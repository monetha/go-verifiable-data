package main

import (
	"flag"
	"fmt"
	"io"
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

var (
	factTypes  = make(map[string]data.Type)
	factSetStr string
)

func init() {
	keys := make([]string, 0, len(data.TypeValues()))

	for _, key := range data.TypeValues() {
		keyStr := strings.ToLower(key.String())
		factTypes[keyStr] = key
		keys = append(keys, keyStr)
	}

	factSetStr = strings.Join(keys, ", ")
}

func main() {
	var (
		backendURL   = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr = cmdutils.AddressVar("passportaddr", common.Address{}, "Ethereum address of passport contract")
		txHash       = cmdutils.HashVar("txhash", common.Hash{}, "the transaction hash to read history value from")
		factTypeStr  = flag.String("ftype", "", fmt.Sprintf("the data type of fact (%v)", factSetStr))
		fileName     = flag.String("out", "", "save retrieved data to the specified file")
		ipfsURL      = flag.String("ipfsurl", "https://ipfs.infura.io:5001", "IPFS node address")
		verbosity    = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule      = flag.String("vmodule", "", "log verbosity pattern")

		err           error
		factType      data.Type
		knownFactType bool
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
	case *fileName == "":
		utils.Fatalf("Use -out to save retrieved data to the specified file")
	case txHash.IsSet() != (*factTypeStr != ""):
		utils.Fatalf("Provide both -txhash and -ftype values")
	case txHash.IsSet():
		if factType, knownFactType = factTypes[*factTypeStr]; !knownFactType {
			utils.Fatalf("Unsupported data type of fact '%v', use one of: %v", *factTypeStr, factSetStr)
		}
	}

	passportAddress := passportAddr.GetValue()
	log.Warn("Loaded configuration", "backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		monethaKey, err := crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232030")
		cmdutils.CheckErr(err, "generating key")
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		passportOwnerKey, err := crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232031")
		cmdutils.CheckErr(err, "generating key")
		passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

		factProviderKey, err := crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
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
		res, err := deployer.New(monethaSession).DeployPassportFactory(ctx)
		cmdutils.CheckErr(err, "create passport factory")
		passportFactoryAddress := res.PassportFactoryAddress

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
		cmdutils.CheckErr(ignoreHash(factProvider.WriteIPFSHash(ctx, passportAddress, factKey, "QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j")), "WriteIPFSHash")

		// some deletes
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteTxData(ctx, passportAddress, factKey)), "DeleteTxData")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteString(ctx, passportAddress, factKey)), "DeleteString")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteBytes(ctx, passportAddress, factKey)), "DeleteBytes")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteAddress(ctx, passportAddress, factKey)), "DeleteAddress")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteUint(ctx, passportAddress, factKey)), "DeleteUint")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteInt(ctx, passportAddress, factKey)), "DeleteInt")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteBool(ctx, passportAddress, factKey)), "DeleteBool")
		cmdutils.CheckErr(ignoreHash(factProvider.DeleteIPFSHash(ctx, passportAddress, factKey)), "DeleteIPFSHash")
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

	if txHash.IsSet() {
		// read history value from transaction
		var fileOp fileOperation

		switch factType {
		case data.TxData:
			hi, err := historian.GetHistoryItemOfWriteTxData(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteTxData")
			fileOp = writeBytes(hi.Data)
		case data.String:
			hi, err := historian.GetHistoryItemOfWriteString(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteString")
			fileOp = writeString(hi.Data)
		case data.Bytes:
			hi, err := historian.GetHistoryItemOfWriteBytes(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteBytes")
			fileOp = writeBytes(hi.Data)
		case data.Address:
			hi, err := historian.GetHistoryItemOfWriteAddress(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteAddress")
			fileOp = writeString(hi.Data.String())
		case data.Uint:
			hi, err := historian.GetHistoryItemOfWriteUint(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteUint")
			fileOp = writeString(hi.Data.String())
		case data.Int:
			hi, err := historian.GetHistoryItemOfWriteInt(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteInt")
			fileOp = writeString(hi.Data.String())
		case data.Bool:
			hi, err := historian.GetHistoryItemOfWriteBool(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "GetHistoryItemOfWriteBool")
			fileOp = writeString(strconv.FormatBool(hi.Data))
		case data.IPFS:
			r, err := facts.NewIPFSDataReader(e, *ipfsURL)
			cmdutils.CheckErr(err, "facts.NewIPFSDataReader")

			rc, err := r.ReadHistoryIPFSData(ctx, passportAddress, txHash.GetValue())
			cmdutils.CheckErr(err, "IPFSDataReader.ReadHistoryIPFSData")

			fileOp = writeReader(rc)
		default:
			cmdutils.CheckErr(fmt.Errorf("unsupported fact type: %v", factType.String()), "reading by type")
		}

		cmdutils.CheckErr(fileOp(f), "Write to file")
	} else {
		// read the whole history
		_, err = f.WriteString("fact_provider,key,data_type,change_type,block_number,tx_hash\n")
		cmdutils.CheckErr(err, "WriteString to file")
		for it.Next() {
			cmdutils.CheckErr(it.Error(), "getting next history item")

			ch := it.Change
			_, err = f.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", ch.FactProvider.Hex(), string(ch.Key[:]), ch.DataType, ch.ChangeType, ch.Raw.BlockNumber, ch.Raw.TxHash.Hex()))
			cmdutils.CheckErr(err, "WriteString to file")
		}
	}

	log.Warn("Done.")
}

func ignoreHash(_ common.Hash, err error) error {
	return err
}

type fileOperation func(f *os.File) error

func writeBytes(b []byte) fileOperation {
	return func(f *os.File) error {
		_, err := f.Write(b)
		return err
	}
}

func writeString(s string) fileOperation {
	return func(f *os.File) error {
		_, err := f.WriteString(s)
		return err
	}
}

func writeReader(r io.Reader) fileOperation {
	return func(f *os.File) error {
		_, err := io.Copy(f, r)
		if rc, ok := r.(io.Closer); ok {
			_ = rc.Close()
		}
		return err
	}
}
