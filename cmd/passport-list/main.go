package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"

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
	"github.com/monetha/reputation-go-sdk/passfactory"
)

func main() {
	var (
		backendURL  = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		factoryAddr = flag.String("factoryaddr", "", "Ethereum address of passport factory contract")
		fileName    = flag.String("out", "", "save retrieved passports to the specified file")
		verbosity   = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule     = flag.String("vmodule", "", "log verbosity pattern")
	)
	flag.Parse()

	if cmd.PrintVersion() {
		return
	}

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	_ = glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *factoryAddr == "" && *backendURL != "":
		utils.Fatalf("Use -factoryaddr to specify an address of passport factory contract")
	case *fileName == "":
		utils.Fatalf("Use -out to save retrieved passports to the specified file")
	}

	passportFactoryAddress := common.HexToAddress(*factoryAddr)
	log.Warn("Loaded configuration", "factory_provider", passportFactoryAddress.Hex(), "backend_url", *backendURL)

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
		passportFactoryAddress, err = deployer.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")

		// creating passport owner session and checking balance
		passportOwnerSession := e.NewSession(passportOwnerKey)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deployer.PassportGasLimit)

		// deploying passport
		passportAddress, err := deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "create passport")

		_ = passportAddress
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		e = eth.New(client, log.Warn)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	pfr := passfactory.NewReader(e)
	filterOpts := &passfactory.PassportFilterOpts{
		Context: ctx,
	}

	it, err := pfr.FilterPassports(filterOpts, passportFactoryAddress)
	cmdutils.CheckErr(err, "FilterPassports")
	defer func() { _ = it.Close() }()

	log.Warn("Writing collected passports to file")

	f, err := os.Create(*fileName)
	cmdutils.CheckErr(err, "Create file")
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("passport_address,first_owner,block_number,tx_hash\n")
	cmdutils.CheckErr(err, "WriteString to file")
	for it.Next() {
		cmdutils.CheckErr(it.Error(), "getting next passport")

		p := it.Passport
		_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v\n", p.ContractAddress.Hex(), p.FirstOwner.Hex(), p.Raw.BlockNumber, p.Raw.TxHash.Hex()))
		cmdutils.CheckErr(err, "WriteString to file")
	}

	log.Warn("Done.")
}
