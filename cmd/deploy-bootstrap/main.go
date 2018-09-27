package main

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"os"

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
)

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		ownerKeyFile = flag.String("ownerkey", "", "owner private key filename")
		ownerKeyHex  = flag.String("ownerkeyhex", "", "private key as hex (for testing)")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		ownerKey *ecdsa.PrivateKey
		err      error
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case *ownerKeyFile == "" && *ownerKeyHex == "":
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key")
	case *ownerKeyFile != "" && *ownerKeyHex != "":
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case *ownerKeyFile != "":
		if ownerKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
			utils.Fatalf("-ownerkey: %v", err)
		}
	case *ownerKeyHex != "":
		if ownerKey, err = crypto.HexToECDSA(*ownerKeyHex); err != nil {
			utils.Fatalf("-ownerkeyhex: %v", err)
		}
	}

	ctx := cmdutils.CreateCtrlCContext()

	ownerAddress := bind.NewKeyedTransactor(ownerKey).From
	log.Warn("Loaded configuration", "owner_address", ownerAddress.Hex(), "backend_url", *backendURL)

	var (
		b backend.Backend
	)
	if *backendURL == "" {
		alloc := core.GenesisAlloc{
			ownerAddress: {Balance: big.NewInt(deployer.PassportFactoryGasLimit)},
		}
		sim := backend.NewSimulatedBackendExtended(alloc, 10000000)
		sim.Commit()

		b = sim
	} else {
		client, err := ethclient.Dial(*backendURL)
		cmdutils.CheckErr(err, "ethclient.Dial")

		b = client
	}

	e := eth.New(b, log.Warn).NewHandleNonceBackend([]common.Address{ownerAddress})
	cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")

	// creating owner session and checking balance
	ownerSession := e.NewSession(ownerKey)
	cmdutils.CheckBalance(ctx, ownerSession, deployer.PassportFactoryGasLimit)

	// deploying passport factory
	_, err = deployer.New(ownerSession).
		DeployPassportFactory(ctx)
	cmdutils.CheckErr(err, "create passport factory")

	log.Warn("Done.")
}
