package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	flag2 "github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/data"
	"github.com/pkg/errors"
)

// ReadFactTxCommand reads passport fact value from transaction
type ReadFactTxCommand struct {
	cmdutils.QuorumPrivateTxICommand
	BackendURL       string                        `long:"backendurl" required:"true" description:"Ethereum backend URL"`
	PassportAddress  flag.EthereumAddress          `long:"passaddr"   required:"true" description:"Ethereum address of passport contract"`
	TXHash           flag2.TXHash                  `long:"txhash"     required:"true" description:"the transaction hash to read value from"`
	FactType         flag2.DataType                `long:"ftype"      required:"true" description:"the data type of fact (txhash, string, bytes, address, uint, int, bool, ipfs, privatedata)"`
	FileName         string                        `long:"out"        required:"true" description:"save fact value to the specified file"`
	IPFSURL          string                        `long:"ipfsurl"                    description:"IPFS node address (to read IPFS and private facts)" default:"https://ipfs.infura.io:5001"`
	PassportOwnerKey *flag.ECDSAPrivateKeyFromFile `long:"ownerkey"                   description:"digital identity owner private key filename (only for privatedata data type)"`
	DataKeyFileName  string                        `long:"datakey"                    description:"data decryption key file name (only for privatedata data type)"`
	Verbosity        int                           `long:"verbosity"                  description:"log verbosity (0-9)" default:"2"`
	VModule          string                        `long:"vmodule"                    description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadFactTxCommand) Execute(args []string) (err error) {
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
		"passport", c.PassportAddress.AsCommonAddress().Hex())

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "cannot create IPFS client")
	}

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, nil)
	if err != nil {
		return err
	}

	historian := facts.NewHistorian(e)

	f, err := os.Create(c.FileName)
	if err != nil {
		return errors.Wrap(err, "could not create output file")
	}
	defer func() { _ = f.Close() }()

	// Read value from transaction
	var fileOp fileOperation

	switch c.FactType.AsType() {
	case data.TxData:
		hi, err := historian.GetHistoryItemOfWriteTxData(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read txdata fact")
		}
		fileOp = writeBytes(hi.Data)
	case data.String:
		hi, err := historian.GetHistoryItemOfWriteString(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read string fact")
		}
		fileOp = writeString(hi.Data)
	case data.Bytes:
		hi, err := historian.GetHistoryItemOfWriteBytes(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read bytes fact")
		}
		fileOp = writeBytes(hi.Data)
	case data.Address:
		hi, err := historian.GetHistoryItemOfWriteAddress(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read address fact")
		}
		fileOp = writeString(hi.Data.String())
	case data.Uint:
		hi, err := historian.GetHistoryItemOfWriteUint(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read uint fact")
		}
		fileOp = writeString(hi.Data.String())
	case data.Int:
		hi, err := historian.GetHistoryItemOfWriteInt(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read int fact")
		}
		fileOp = writeString(hi.Data.String())
	case data.Bool:
		hi, err := historian.GetHistoryItemOfWriteBool(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read bool fact")
		}
		fileOp = writeString(strconv.FormatBool(hi.Data))
	case data.IPFS:
		r := facts.NewIPFSDataReader(e, fs)

		rc, err := r.ReadHistoryIPFSData(ctx, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
		if err != nil {
			return errors.Wrap(err, "could not read IPFS fact")
		}

		fileOp = writeReader(rc)
	case data.PrivateData:
		rd := facts.NewPrivateDataReader(e, fs)

		if c.PassportOwnerKey != nil {
			decryptedData, err := rd.ReadHistoryPrivateData(ctx, c.PassportOwnerKey.AsECDSAPrivateKey(), c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
			if err != nil {
				return errors.Wrap(err, "could not read privatedata fact using owner key")
			}
			fileOp = writeBytes(decryptedData)
		} else if privateDataSecretKey != nil {
			decryptedData, err := rd.ReadHistoryPrivateDataUsingSecretKey(ctx, privateDataSecretKey, c.PassportAddress.AsCommonAddress(), c.TXHash.AsCommonHash())
			if err != nil {
				return errors.Wrap(err, "could not read privatedata fact using secret key")
			}
			fileOp = writeBytes(decryptedData)
		} else {
			return fmt.Errorf("either specify owner key or secret key")
		}
	default:
		return fmt.Errorf("unsupported fact type: %v", c.FactType.AsType().String())
	}

	if err := fileOp(f); err != nil {
		return errors.Wrap(err, "could not write to file")
	}

	log.Warn("Done.")
	return nil
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
