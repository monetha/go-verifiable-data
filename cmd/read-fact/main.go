package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/deployer"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/data"
)

func main() {
	var (
		backendURL       = flag.String("backendurl", "", "backend URL (simulated backend used if empty)")
		passportAddr     = cmdutils.AddressVar("passportaddr", common.Address{}, "Ethereum address of passport contract")
		factProviderAddr = cmdutils.AddressVar("factprovideraddr", common.Address{}, "Ethereum address of fact provider")
		factKeyStr       = flag.String("fkey", "", "the key of the fact (max. 32 bytes)")
		factTypeVar      = cmdutils.DataTypeFlagVar("ftype", data.TxData, fmt.Sprintf("the data type of fact (%v)", cmdutils.DataTypeSetStr()))
		fileName         = flag.String("out", "", "save retrieved data to the specified file")
		ownerKeyFile     = flag.String("ownerkey", "", "passport owner private key filename (only for privatedata data type)")
		ownerKeyHex      = cmdutils.PrivateKeyFlagVar("ownerkeyhex", nil, "passport owner private key as hex (only for privatedata data type)")
		dataKeyFileName  = flag.String("datakeyfile", "", "data decryption key file name (only for privatedata data type)")
		ipfsURL          = flag.String("ipfsurl", "https://ipfs.infura.io:5001", "IPFS node address")
		verbosity        = flag.Int("verbosity", int(log.LvlWarn), "log verbosity (0-9)")
		vmodule          = flag.String("vmodule", "", "log verbosity pattern")
		quorumEnclave    = flag.String("quorum_enclave", "", "Quorum enclave url to decrypt facts, stored using private transactions")

		factKey [32]byte
		err     error
		// private data variables
		passportOwnerKey     *ecdsa.PrivateKey
		privateDataSecretKey []byte
	)
	flag.Parse()

	if cmd.HasPrintedVersion() {
		return
	}

	bf := cmdutils.NewBackendFactory(quorumEnclave, nil)

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	_ = glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	switch {
	case !passportAddr.IsSet() && *backendURL != "":
		utils.Fatalf("Use -passportaddr to specify an address of passport contract")
	case !factProviderAddr.IsSet() && *backendURL != "":
		utils.Fatalf("Use -factprovideraddr to specify fact provider address")
	case *fileName == "":
		utils.Fatalf("Use -out to save retrieved data to the specified file")
	case *factKeyStr == "":
		utils.Fatalf("Use -fkey to specify the key of the fact")
	case !factTypeVar.IsSet():
		utils.Fatalf("Use -ftype to specify the data type of fact")
	case *ownerKeyFile != "" && factTypeVar.GetValue() != data.PrivateData:
		utils.Fatalf("Use -ownerkey only with -ftype privatedata")
	case ownerKeyHex.IsSet() && factTypeVar.GetValue() != data.PrivateData:
		utils.Fatalf("Use -ownerkeyhex only with -ftype privatedata")
	case *dataKeyFileName != "" && factTypeVar.GetValue() != data.PrivateData:
		utils.Fatalf("Use -datakeyfile only with -ftype privatedata")
	case *ownerKeyFile == "" && !ownerKeyHex.IsSet() && *dataKeyFileName == "" && factTypeVar.GetValue() == data.PrivateData:
		utils.Fatalf("Use -ownerkey or -ownerkeyhex to specify a private key of passport owner or -datakeyfile to provide decryption key")
	case *ownerKeyFile != "" && ownerKeyHex.IsSet():
		utils.Fatalf("Options -ownerkey or -ownerkeyhex are mutually exclusive")
	case *ownerKeyFile != "" && *dataKeyFileName != "":
		utils.Fatalf("Options -ownerkey or -datakeyfile are mutually exclusive")
	case ownerKeyHex.IsSet() && *dataKeyFileName != "":
		utils.Fatalf("Options -ownerkeyhex or -datakeyfile are mutually exclusive")
	}

	if factKeyBytes := []byte(*factKeyStr); len(factKeyBytes) <= 32 {
		copy(factKey[:], factKeyBytes)
	} else {
		utils.Fatalf("The key string should fit into 32 bytes")
	}

	if factTypeVar.GetValue() == data.PrivateData {
		switch {
		case *dataKeyFileName != "":
			if privateDataSecretKey, err = ioutil.ReadFile(*dataKeyFileName); err != nil {
				utils.Fatalf("-datakeyfile: %v", err)
			}
		case *ownerKeyFile != "":
			if passportOwnerKey, err = crypto.LoadECDSA(*ownerKeyFile); err != nil {
				utils.Fatalf("-ownerkey: %v", err)
			}
		case ownerKeyHex.IsSet():
			passportOwnerKey = ownerKeyHex.GetValue()
		}
	}

	passportAddress := passportAddr.GetValue()
	factProviderAddress := factProviderAddr.GetValue()
	log.Warn("Loaded configuration", "fact_provider", factProviderAddress.Hex(), "fact_key", factKey,
		"backend_url", *backendURL, "passport", passportAddress.Hex())

	ctx := cmdutils.CreateCtrlCContext()

	fs, err := ipfs.New(*ipfsURL)
	cmdutils.CheckErr(err, "creating IPFS client")

	var (
		e *eth.Eth
	)
	if *backendURL == "" {
		// use deterministic "random" numbers in simulated environment
		randReader := cmdutils.NewMathRandFixedSeed()

		monethaKey := cmdutils.TestMonethaAdminKey
		monethaAddress := bind.NewKeyedTransactor(monethaKey).From

		passportOwnerKey := cmdutils.TestPassportOwnerKey
		passportOwnerAddress := bind.NewKeyedTransactor(passportOwnerKey).From

		factProviderKey := cmdutils.TestFactProviderKey
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
		res, err := deployer.New(monethaSession).Bootstrap(ctx)
		cmdutils.CheckErr(err, "create passport factory")
		passportFactoryAddress := res.PassportFactoryAddress

		// creating passport owner session and checking balance
		passportOwnerSession := e.NewSession(passportOwnerKey)
		cmdutils.CheckBalance(ctx, passportOwnerSession, deployer.PassportGasLimit)

		// deploying passport
		passportAddress, err = deployer.New(passportOwnerSession).DeployPassport(ctx, passportFactoryAddress)
		cmdutils.CheckErr(err, "create passport")

		factProviderSession := e.NewSession(factProviderKey)
		factProvider := facts.NewProvider(factProviderSession)

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

		_, err = factProvider.WriteIPFSHash(ctx, passportAddress, factKey, "QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j")
		cmdutils.CheckErr(err, "WriteIPFSHash")

		wr := facts.NewPrivateDataWriter(factProviderSession, fs)

		wpdRes, err := wr.WritePrivateData(ctx, passportAddress, factKey, []byte("this is a secret message"), randReader)
		cmdutils.CheckErr(err, "writing private data")
		log.Warn("Private data has been written", "data_secret_key", wpdRes.DataKey)
	} else {
		client, err := bf.DialBackend(*backendURL)
		cmdutils.CheckErr(err, "bf.DialBackend")

		e = bf.NewEth(ctx, client)
		cmdutils.CheckErr(e.UpdateSuggestedGasPrice(ctx), "SuggestGasPrice")
	}

	factReader := facts.NewReader(e)
	var fileOp fileOperation

	factType := factTypeVar.GetValue()
	switch factType {
	case data.TxData:
		dataBytes, err := factReader.ReadTxData(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadTxData")
		fileOp = writeBytes(dataBytes)
	case data.String:
		dataBytes, err := factReader.ReadString(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadString")
		fileOp = writeString(dataBytes)
	case data.Bytes:
		dataBytes, err := factReader.ReadBytes(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadBytes")
		fileOp = writeBytes(dataBytes)
	case data.Address:
		addressVal, err := factReader.ReadAddress(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadAddress")
		fileOp = writeString(addressVal.String())
	case data.Uint:
		uintVal, err := factReader.ReadUint(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadUint")
		fileOp = writeString(uintVal.String())
	case data.Int:
		intVal, err := factReader.ReadInt(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadInt")
		fileOp = writeString(intVal.String())
	case data.Bool:
		dataBytes, err := factReader.ReadBool(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "ReadBool")
		fileOp = writeString(strconv.FormatBool(dataBytes))
	case data.IPFS:
		r := facts.NewIPFSDataReader(e, fs)

		rc, err := r.ReadIPFSData(ctx, passportAddress, factProviderAddress, factKey)
		cmdutils.CheckErr(err, "IPFSDataReader.ReadIPFSData")

		fileOp = writeReader(rc)
	case data.PrivateData:
		rd := facts.NewPrivateDataReader(e, fs)

		if passportOwnerKey != nil {
			decryptedData, err := rd.ReadPrivateData(ctx, passportOwnerKey, passportAddress, factProviderAddress, factKey)
			cmdutils.CheckErr(err, "ReadPrivateData")
			fileOp = writeBytes(decryptedData)
		} else if privateDataSecretKey != nil {
			decryptedData, err := rd.ReadPrivateDataUsingSecretKey(ctx, privateDataSecretKey, passportAddress, factProviderAddress, factKey)
			cmdutils.CheckErr(err, "ReadPrivateDataUsingSecretKey")
			fileOp = writeBytes(decryptedData)
		} else {
			utils.Fatalf("either specify passport owner secret key or data decryption secret key")
		}
	default:
		cmdutils.CheckErr(fmt.Errorf("unsupported fact type: %v", factType.String()), "reading by type")
	}

	log.Warn("Writing data to file")

	f, err := os.Create(*fileName)
	cmdutils.CheckErr(err, "Create file")
	defer func() { _ = f.Close() }()

	cmdutils.CheckErr(fileOp(f), "Write to file")

	log.Warn("Done.")
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
