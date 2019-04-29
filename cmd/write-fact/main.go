package main

import (
	"bufio"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/reputation-go-sdk/cmd"
	"github.com/monetha/reputation-go-sdk/cmd/internal/cmdutils"
	"github.com/monetha/reputation-go-sdk/deployer"
	"github.com/monetha/reputation-go-sdk/eth"
	"github.com/monetha/reputation-go-sdk/eth/backend"
	"github.com/monetha/reputation-go-sdk/facts"
	"github.com/monetha/reputation-go-sdk/types/data"
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr = cmdutils.AddressVar("passportaddr", common.Address{}, "Ethereum address of passport contract")
		factKeyStr   = flag.String("fkey", "", "the key of the fact (max. 32 bytes)")
		factTypeVar  = cmdutils.DataTypeFlagVar("ftype", data.TxData, fmt.Sprintf("the data type of fact (%v)", cmdutils.DataTypeSetStr()))
		ownerKeyFile = flag.String("ownerkey", "", "fact provider private key filename")
		ownerKeyHex  = cmdutils.PrivateKeyFlagVar("ownerkeyhex", nil, "fact provider private key as hex")
		ipfsURL      = flag.String("ipfsurl", "https://ipfs.infura.io:5001", "IPFS node address")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		factProviderKey *ecdsa.PrivateKey
		factKey         [32]byte
		factBytes       []byte
		factString      string
		factAddress     common.Address
		factInt         *big.Int
		factBool        bool
		err             error
	)
	flag.Parse()

	if cmd.PrintVersion() {
		return
	}

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case !passportAddr.IsSet() && *backendURL != "":
		utils.Fatalf("Use -passportaddr to specify an address of passport contract")
	case *factKeyStr == "":
		utils.Fatalf("Use -fkey to specify the key of the fact")
	case !factTypeVar.IsSet():
		utils.Fatalf("Use -ftype to specify the data type of fact")
	case *ownerKeyFile == "" && !ownerKeyHex.IsSet():
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key of fact provider")
	case *ownerKeyFile != "" && ownerKeyHex.IsSet():
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case *ownerKeyFile != "":
		if factProviderKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
			utils.Fatalf("-ownerkey: %v", err)
		}
	case ownerKeyHex.IsSet():
		factProviderKey = ownerKeyHex.GetValue()
	}

	if factKeyBytes := []byte(*factKeyStr); len(factKeyBytes) <= 32 {
		copy(factKey[:], factKeyBytes)
	} else {
		utils.Fatalf("The key string should fit into 32 bytes")
	}

	// parse fact data
	factType := factTypeVar.GetValue()
	switch {
	case factType == data.TxData || factType == data.Bytes:
		if factBytes, err = ioutil.ReadAll(os.Stdin); err != nil {
			utils.Fatalf("failed to read fact bytes: %v", err)
		}
	case factType == data.String:
		if factString, err = copyToString(os.Stdin); err != nil {
			utils.Fatalf("failed to read fact string: %v", err)
		}
	case factType == data.Address:
		factAddress = parseAddress(os.Stdin)
	case factType == data.Uint:
		factInt = parseBigInt(os.Stdin)
		if factInt.Cmp(new(big.Int)) == -1 {
			utils.Fatalf("expected non-negative number, but got %v", factInt)
		}
	case factType == data.Int:
		factInt = parseBigInt(os.Stdin)
	case factType == data.Bool:
		var boolStr string
		if boolStr, err = readLine(os.Stdin); err != nil {
			utils.Fatalf("failed to read fact bool: %v", err)
		}
		if factBool, err = strconv.ParseBool(boolStr); err != nil {
			utils.Fatalf("invalid fact bool: %v", boolStr)
		}
	}

	passportAddress := passportAddr.GetValue()
	factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From
	log.Warn("Loaded configuration", "fact_provider", factProviderAddress.Hex(), "backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		monethaKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		passportOwnerKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

		alloc := core.GenesisAlloc{
			monethaAddress:       {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
			passportOwnerAddress: {Balance: big.NewInt(deployer.PassportGasLimit)},
			factProviderAddress:  {Balance: big.NewInt(10000000000000)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		e = eth.New(sim, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")

		// creating owner session and checking balance
		monethaSession := e.NewSession(monethaKey)
		cmdutils.CheckBalance(ctx, monethaSession, deployer.PassportFactoryGasLimit)

		// deploying passport factory
		res, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")
		passportFactoryAddress := res.PassportFactoryAddress

		// creating passport owner session and checking balance
		passportOwnerSession := e.NewSession(passportOwnerKey)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deployer.PassportGasLimit)

		// deploying passport
		passportAddress, err = deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "create passport")
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	factProviderSession := e.NewSession(factProviderKey)

	// TODO: check balance

	provider := facts.NewProvider(factProviderSession)
	switch factType {
	case data.TxData:
		cmdutils.CheckErr(ignoreHash(provider.WriteTxData(ctx, passportAddress, factKey, factBytes)), "WriteTxData")
	case data.String:
		cmdutils.CheckErr(ignoreHash(provider.WriteString(ctx, passportAddress, factKey, factString)), "WriteString")
	case data.Bytes:
		cmdutils.CheckErr(ignoreHash(provider.WriteBytes(ctx, passportAddress, factKey, factBytes)), "WriteBytes")
	case data.Address:
		cmdutils.CheckErr(ignoreHash(provider.WriteAddress(ctx, passportAddress, factKey, factAddress)), "WriteAddress")
	case data.Uint:
		cmdutils.CheckErr(ignoreHash(provider.WriteUint(ctx, passportAddress, factKey, factInt)), "WriteUint")
	case data.Int:
		cmdutils.CheckErr(ignoreHash(provider.WriteInt(ctx, passportAddress, factKey, factInt)), "WriteInt")
	case data.Bool:
		cmdutils.CheckErr(ignoreHash(provider.WriteBool(ctx, passportAddress, factKey, factBool)), "WriteBool")
	case data.IPFS:
		log.Warn("Uploading data to IPFS...", "url", *ipfsURL)
		w, err := facts.NewIPFSDataWriter(factProviderSession, *ipfsURL)
		cmdutils.CheckErr(err, "facts.NewIPFSDataWriter")

		_, err = w.WriteIPFSData(ctx, passportAddress, factKey, os.Stdin)
		cmdutils.CheckErr(err, "writing IPFS data")
	default:
		cmdutils.CheckErr(fmt.Errorf("unsupported fact type: %v", factType.String()), "writing by type")
	}

	log.Warn("Done.")
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}

func copyToString(r io.Reader) (res string, err error) {
	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}

func readLine(r io.Reader) (res string, err error) {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		res = scanner.Text()
	} else {
		err = scanner.Err()
	}
	return
}

func parseAddress(r io.Reader) common.Address {
	var (
		addressStr string
		err        error
	)
	if addressStr, err = readLine(r); err != nil {
		utils.Fatalf("failed to read fact address: %v", err)
	}
	if !common.IsHexAddress(addressStr) {
		utils.Fatalf("invalid fact address: %v", addressStr)
	}
	return common.HexToAddress(addressStr)
}

func parseBigInt(r io.Reader) (res *big.Int) {
	var (
		s   string
		err error
	)
	if s, err = readLine(r); err != nil {
		utils.Fatalf("failed to read fact number: %v", err)
	}
	var ok bool
	if res, ok = new(big.Int).SetString(s, 0); !ok {
		utils.Fatalf("failed to parse number: %v", s)
	}
	return
}
