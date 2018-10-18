package main

import (
	"bytes"
	"flag"
	"fmt"
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
	"gitlab.com/monetha/protocol-go-sdk/cmd/internal/cmdutils"
	"gitlab.com/monetha/protocol-go-sdk/deployer"
	"gitlab.com/monetha/protocol-go-sdk/eth"
	"gitlab.com/monetha/protocol-go-sdk/eth/backend"
	"gitlab.com/monetha/protocol-go-sdk/facts"
)

type factType int

const (
	ftTxData  factType = iota
	ftString  factType = iota
	ftBytes   factType = iota
	ftAddress factType = iota
	ftUint    factType = iota
	ftInt     factType = iota
	ftBool    factType = iota
)

var (
	factTypes = map[string]factType{
		"txdata":  ftTxData,
		"string":  ftString,
		"bytes":   ftBytes,
		"address": ftAddress,
		"uint":    ftUint,
		"int":     ftInt,
		"bool":    ftBool,
	}

	factSetStr string
)

func init() {
	keys := make([]string, 0, len(factTypes))
	for key := range factTypes {
		keys = append(keys, key)
	}

	factSetStr = strings.Join(keys, ", ")
}

func main() {
	var (
		backendURL       = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr     = flag.String("passportaddr", "", "Ethereum address of passport contract")
		factProviderAddr = flag.String("factprovideraddr", "", "Ethereum address of fact provider")
		factKeyStr       = flag.String("fkey", "", "the key of the fact (max. 32 bytes)")
		factTypeStr      = flag.String("ftype", "", fmt.Sprintf("the data type of fact (%v)", factSetStr))
		fileName         = flag.String("out", "", "save retrieved data to the specified file")
		verbosity        = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule          = flag.String("vmodule", "", "log verbosity pattern")

		factKey       [32]byte
		factType      factType
		knownFactType bool
		err           error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	factType, knownFactType = factTypes[*factTypeStr]

	switch {
	case *passportAddr == "" && *backendURL != "":
		utils.Fatalf("Use -passportaddr to specify an address of passport contract")
	case *factProviderAddr == "" && *backendURL != "":
		utils.Fatalf("Use -factprovideraddr to specify fact provider address")
	case *fileName == "":
		utils.Fatalf("Use -out to save retrieved data to the specified file")
	case *factKeyStr == "":
		utils.Fatalf("Use -fkey to specify the key of the fact")
	case *factTypeStr == "":
		utils.Fatalf("Use -ftype to specify the data type of fact")
	case !knownFactType:
		utils.Fatalf("Unsupported data type of fact '%v', use one of: %v", *factTypeStr, factSetStr)
	}

	if factKeyBytes := []byte(*factKeyStr); len(factKeyBytes) <= 32 {
		copy(factKey[:], factKeyBytes)
	} else {
		utils.Fatalf("The key string should fit into 32 bytes")
	}

	passportAddress := common.HexToAddress(*passportAddr)
	factProviderAddress := common.HexToAddress(*factProviderAddr)
	log.Warn("Loaded configuration", "fact_provider", factProviderAddress.Hex(), "fact_key", factKey,
		"backend_url", *backendURL, "passport", passportAddress.Hex())

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

		factProviderKey, err := crypto.GenerateKey()
		cmdutils.CheckErr(err, "generating key")
		factProviderAddress = bind.NewKeyedTransactor(factProviderKey).From

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
		passportFactoryAddress, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")

		// creating passport owner session and checking balance
		passportOwnerSession := e.NewSession(passportOwnerKey)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deployer.PassportGasLimit)

		// deploying passport
		passportAddress, err = deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "create passport")

		factProvider := facts.NewProvider(e.NewSession(factProviderKey))

		_, err = factProvider.WriteTxData(ctx, passportAddress, factKey, []byte("This is test only tx data!"))
		cmdutils.CheckErr(err, "WriteTxData")

		_, err = factProvider.WriteString(ctx, passportAddress, factKey, "This is test only string data!")
		cmdutils.CheckErr(err, "WriteString")

		_, err = factProvider.WriteBytes(ctx, passportAddress, factKey, []byte("This is test only bytes data!"))
		cmdutils.CheckErr(err, "WriteBytes")

		_, err = factProvider.WriteAddress(ctx, passportAddress, factKey, factProviderAddress)
		cmdutils.CheckErr(err, "WriteAddress")

		_, err = factProvider.WriteUint(ctx, passportAddress, factKey, big.NewInt(123456789))
		cmdutils.CheckErr(err, "WriteUint")

		_, err = factProvider.WriteInt(ctx, passportAddress, factKey, big.NewInt(987654321))
		cmdutils.CheckErr(err, "WriteInt")

		_, err = factProvider.WriteBool(ctx, passportAddress, factKey, true)
		cmdutils.CheckErr(err, "WriteBool")
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	factReader := facts.NewReader(e)
	buf := new(bytes.Buffer)

	switch factType {
	case ftTxData:
		data, err := factReader.ReadTxData(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadTxData")
		buf.Write(data)
	case ftString:
		data, err := factReader.ReadString(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadString")
		buf.WriteString(data)
	case ftBytes:
		data, err := factReader.ReadBytes(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadBytes")
		buf.Write(data)
	case ftAddress:
		data, err := factReader.ReadAddress(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadAddress")
		buf.WriteString(data.String())
	case ftUint:
		data, err := factReader.ReadUint(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadUint")
		buf.WriteString(data.String())
	case ftInt:
		data, err := factReader.ReadInt(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadInt")
		buf.WriteString(data.String())
	case ftBool:
		data, err := factReader.ReadBool(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadBool")
		buf.WriteString(strconv.FormatBool(data))
	}

	log.Warn("Writing data to file")

	f, err := os.Create(*fileName)
	cmdutils.CheckErr(err, "Create file")
	defer f.Close()

	_, err = buf.WriteTo(f)
	cmdutils.CheckErr(err, "Write to file")

	log.Warn("Done.")
}
