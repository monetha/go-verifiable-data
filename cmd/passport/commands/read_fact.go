package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/data"
	"github.com/pkg/errors"
)

// ReadFactCommand reads latest passport fact value
type ReadFactCommand struct {
	cmdutils.QuorumPrivateTxICommand
	BackendURL          string                `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	PassportAddress     flag.EthereumAddress `long:"passaddr"         required:"true" description:"Ethereum address of passport contract"`
	FactProviderAddress flag.EthereumAddress `long:"factprovideraddr" required:"true" description:"Ethereum address of fact provider"`
	FactKey             flag.FactKey         `long:"fkey"             required:"true" description:"the key of the fact"`
	FactType            flag.DataType        `long:"ftype"            required:"true" description:"the data type of fact (txhash, string, bytes, address, uint, int, bool, ipfs, privatedata)"`
	FileName            string                `long:"out"              required:"true" description:"save fact value to the specified file"`
	IPFSURL             string                `long:"ipfsurl"                          description:"IPFS node address (to read IPFS and private facts)" default:"https://ipfs.infura.io:5001"`
	PassportOwnerKey    *flag.ECDSAPrivateKeyFromFile `long:"ownerkey"                         description:"digital identity owner private key filename (only for privatedata data type)"`
	DataKeyFileName     string                         `long:"datakey"                          description:"data decryption key file name (only for privatedata data type)"`
	Verbosity           int                            `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule             string                         `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadFactCommand) Execute(args []string) (err error) {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	switch {
	case c.DataKeyFileName != "" && c.FactType.AsType() != data.PrivateData:
		return fmt.Errorf("use --datakey only with --ftype privatedata")
	case c.PassportOwnerKey != nil && c.FactType.AsType() != data.PrivateData:
		return fmt.Errorf("use --ownerkey only with --ftype privatedata")
	case c.PassportOwnerKey == nil && c.DataKeyFileName == "" && c.FactType.AsType() == data.PrivateData:
		return fmt.Errorf("use --ownerkey to specify a private key of passport owner or --datakey to provide decryption key")
	case c.PassportOwnerKey != nil && c.DataKeyFileName != "":
		return fmt.Errorf("options --ownerkey or --datakey are mutually exclusive")
	}

	var privateDataSecretKey []byte

	if c.DataKeyFileName != "" {
		if privateDataSecretKey, err = ioutil.ReadFile(c.DataKeyFileName); err != nil {
			return errors.Wrap(err, "cannot read data key file")
		}
	}

	log.Warn("Loaded configuration",
		"backend_url", c.BackendURL,
		"passport", c.PassportAddress.AsCommonAddress().Hex(),
		"fact_provider", c.FactProviderAddress.AsCommonAddress().Hex(),
		"fact_key", c.FactKey.AsString())

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "cannot create IPFS client")
	}

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, nil)
	if err != nil {
		return err
	}

	factReader := facts.NewReader(e)

	// Read value
	var fileOp fileOperation

	switch c.FactType.AsType() {
	case data.TxData:
		data, err := factReader.ReadTxData(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read txdata fact")
		}
		fileOp = writeBytes(data)
	case data.String:
		data, err := factReader.ReadString(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read string fact")
		}
		fileOp = writeString(data)
	case data.Bytes:
		data, err := factReader.ReadBytes(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read bytes fact")
		}
		fileOp = writeBytes(data)
	case data.Address:
		data, err := factReader.ReadAddress(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read address fact")
		}
		fileOp = writeString(data.String())
	case data.Uint:
		data, err := factReader.ReadUint(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read uint fact")
		}
		fileOp = writeString(data.String())
	case data.Int:
		data, err := factReader.ReadInt(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read int fact")
		}
		fileOp = writeString(data.String())
	case data.Bool:
		data, err := factReader.ReadBool(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read bool fact")
		}
		fileOp = writeString(strconv.FormatBool(data))
	case data.IPFS:
		r := facts.NewIPFSDataReader(e, fs)

		rc, err := r.ReadIPFSData(ctx, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
		if err != nil {
			return errors.Wrap(err, "could not read IPFS fact")
		}

		fileOp = writeReader(rc)
	case data.PrivateData:
		rd := facts.NewPrivateDataReader(e, fs)

		if c.PassportOwnerKey != nil {
			decryptedData, err := rd.ReadPrivateData(
				ctx, c.PassportOwnerKey.AsECDSAPrivateKey(), c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
			if err != nil {
				return errors.Wrap(err, "could not read privatedata fact using owner key")
			}
			fileOp = writeBytes(decryptedData)
		} else if privateDataSecretKey != nil {
			decryptedData, err := rd.ReadPrivateDataUsingSecretKey(
				ctx, privateDataSecretKey, c.PassportAddress.AsCommonAddress(), c.FactProviderAddress.AsCommonAddress(), c.FactKey)
			if err != nil {
				return errors.Wrap(err, "could not read privatedata fact using secret data key")
			}
			fileOp = writeBytes(decryptedData)
		} else {
			return fmt.Errorf("either specify owner key or secret key")
		}
	default:
		return fmt.Errorf("unsupported fact type: %v", c.FactType.AsType().String())
	}

	log.Warn("Writing data to file")

	f, err := os.Create(c.FileName)
	if err != nil {
		return errors.Wrap(err, "could not create output file")
	}
	defer func() { _ = f.Close() }()

	if err := fileOp(f); err != nil {
		return errors.Wrap(err, "could not write to file")
	}

	log.Warn("Done.")
	return nil
}
