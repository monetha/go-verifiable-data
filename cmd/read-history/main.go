package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
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
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr = flag.String("passportaddr", "", "Ethereum address of passport contract")
		fileName     = flag.String("out", "", "save retrieved data to the specified file")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		err error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *passportAddr == "" && *backendURL != "":
		utils.Fatalf("Use -passportaddr to specify an address of passport contract")
	case *fileName == "":
		utils.Fatalf("Use -out to save retrieved data to the specified file")
	}

	passportAddress := common.HexToAddress(*passportAddr)
	log.Warn("Loaded configuration", "backend_url", *backendURL, "passport", passportAddress.Hex())

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
		factProviderAddress := bind.NewKeyedTransactor(factProviderKey).From

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

		var factKey [32]byte
		copy(factKey[:], "test_key")

		// some writes
		cmdutils.CheckErr(ignoreHash(factProvider.WriteTxData(ctx, passportAddress, factKey, []byte("This is test only tx data!"))), "WriteTxData")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteString(ctx, passportAddress, factKey, "This is test only string data!")), "WriteString")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteBytes(ctx, passportAddress, factKey, []byte("This is test only bytes data!"))), "WriteBytes")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteAddress(ctx, passportAddress, factKey, factProviderAddress)), "WriteAddress")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteUint(ctx, passportAddress, factKey, big.NewInt(123456789))), "WriteUint")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteInt(ctx, passportAddress, factKey, big.NewInt(-987654321))), "WriteInt")
		cmdutils.CheckErr(ignoreHash(factProvider.WriteBool(ctx, passportAddress, factKey, true)), "WriteBool")

		// some deletes
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteTxData(ctx, passportAddress, factKey)), "DeleteTxData")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteString(ctx, passportAddress, factKey)), "DeleteString")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteBytes(ctx, passportAddress, factKey)), "DeleteBytes")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteAddress(ctx, passportAddress, factKey)), "DeleteAddress")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteUint(ctx, passportAddress, factKey)), "DeleteUint")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteInt(ctx, passportAddress, factKey)), "DeleteInt")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteBool(ctx, passportAddress, factKey)), "DeleteBool")
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	historian := facts.NewHistorian(e)

	it, err := historian.FilterChanges(&facts.ChangesFilterOpts{Context: ctx}, passportAddress)
	cmdutils.CheckErr(err, "FilterChanges")

	f, err := os.Create(*fileName)
	cmdutils.CheckErr(err, "Create file")
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("fact_provider,key,data_type,change_type,block_number,tx_hash\n")
	cmdutils.CheckErr(err, "WriteString to file")
	for it.Next() {
		cmdutils.CheckErr(it.Error(), "getting next history item")

		ch := it.Change
		_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", ch.FactProvider.Hex(), string(ch.Key[:]), ch.DataType, ch.ChangeType, ch.Raw.BlockNumber, ch.Raw.TxHash.Hex()))
		cmdutils.CheckErr(err, "WriteString to file")
	}

	log.Warn("Done.")
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}
