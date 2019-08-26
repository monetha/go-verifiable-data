package commands

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/data"
	"github.com/pkg/errors"
)

// WriteFactCommand writes fact to passport
type WriteFactCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	BackendURL      string                         `long:"backendurl"       required:"true" description:"Ethereum backend URL"`
	PassportAddress flag.EthereumAddress          `long:"passaddr"         required:"true" description:"Ethereum address of passport contract"`
	FactProviderKey *flag.ECDSAPrivateKeyFromFile `long:"factproviderkey"  required:"true" description:"data source (fact provider) private key filename"`
	FactKey         flag.FactKey                  `long:"fkey"             required:"true" description:"the key of the fact (max 32 bytes)"`
	FactType        flag.DataType                 `long:"ftype"            required:"true" description:"the data type of fact (txhash, string, bytes, address, uint, int, bool, ipfs, privatedata)"`
	IPFSURL         string                         `long:"ipfsurl"                          description:"IPFS node address (to write IPFS and private facts)" default:"https://ipfs.infura.io:5001"`
	DataKeyFileName string                         `long:"datakey"                          description:"save data encryption key to the specified file (only for privatedata data type)"`
	Verbosity       int                           `long:"verbosity"                        description:"log verbosity (0-9)" default:"2"`
	VModule         string                        `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *WriteFactCommand) Execute(args []string) (err error) {
	var (
		factBytes   []byte
		factString  string
		factAddress *common.Address
		factInt     *big.Int
		factBool    bool
	)

	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	switch {
	case c.DataKeyFileName != "" && c.FactType.AsType() != data.PrivateData:
		return fmt.Errorf("use --datakey only with --ftype privatedata")
	}

	// Parse fact data
	factType := c.FactType.AsType()
	switch {
	case factType == data.TxData || factType == data.Bytes || factType == data.PrivateData:
		if factBytes, err = ioutil.ReadAll(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact bytes")
		}
	case factType == data.String:
		if factString, err = copyToString(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact string")
		}
	case factType == data.Address:
		if factAddress, err = parseAddress(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact address")
		}
	case factType == data.Uint:
		if factInt, err = parseBigInt(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact uint")
		}
		if factInt.Cmp(new(big.Int)) == -1 {
			return fmt.Errorf("invalid fact uint: expected non-negative number, but got %v", factInt)
		}
	case factType == data.Int:
		if factInt, err = parseBigInt(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact int")
		}
	case factType == data.Bool:
		var boolStr string
		if boolStr, err = readLine(os.Stdin); err != nil {
			return errors.Wrap(err, "could not read fact bool")
		}
		if factBool, err = strconv.ParseBool(boolStr); err != nil {
			return fmt.Errorf("invalid fact bool: %v", boolStr)
		}
	}

	factProviderAddress := bind.NewKeyedTransactor(c.FactProviderKey.AsECDSAPrivateKey()).From
	log.Warn("Loaded configuration",
		"backend_url", c.BackendURL,
		"passport", c.PassportAddress.AsCommonAddress().Hex(),
		"fact_provider", factProviderAddress.Hex(),
		"fact_key", c.FactKey.AsString())

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	e = e.NewHandleNonceBackend([]common.Address{factProviderAddress})

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "cannot create IPFS client")
	}

	randReader := rand.Reader

	factProviderSession := e.NewSession(c.FactProviderKey.AsECDSAPrivateKey())

	provider := facts.NewProvider(factProviderSession)
	switch factType {
	case data.TxData:
		if err = ignoreHash(provider.WriteTxData(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factBytes)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.String:
		if err = ignoreHash(provider.WriteString(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factString)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.Bytes:
		if err = ignoreHash(provider.WriteBytes(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factBytes)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.Address:
		if err = ignoreHash(provider.WriteAddress(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, *factAddress)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.Uint:
		if err = ignoreHash(provider.WriteUint(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factInt)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.Int:
		if err = ignoreHash(provider.WriteInt(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factInt)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.Bool:
		if err = ignoreHash(provider.WriteBool(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factBool)); err != nil {
			return errors.Wrap(err, "could not write fact")
		}
	case data.IPFS:
		log.Warn("Uploading data to IPFS...", "url", c.IPFSURL)
		w := facts.NewIPFSDataWriter(factProviderSession, fs)

		_, err = w.WriteIPFSData(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, os.Stdin)
		if err != nil {
			return errors.Wrap(err, "could not write to IPFS")
		}
	case data.PrivateData:
		wr := facts.NewPrivateDataWriter(factProviderSession, fs)

		res, err := wr.WritePrivateData(ctx, c.PassportAddress.AsCommonAddress(), c.FactKey, factBytes, randReader)
		if err != nil {
			return errors.Wrap(err, "could not write private data")
		}

		if c.DataKeyFileName != "" {
			log.Warn("Writing data encryption key to file", "file_name", c.DataKeyFileName)
			if ioutil.WriteFile(c.DataKeyFileName, res.DataKey, 0400) != nil {
				return errors.Wrap(err, "could not write data encryption key to file")
			}
		}
	default:
		return fmt.Errorf("unsupported fact type: %v", factType.String())
	}

	log.Warn("Done.")
	return nil
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

func parseAddress(r io.Reader) (*common.Address, error) {
	var (
		addressStr string
		err        error
	)

	if addressStr, err = readLine(r); err != nil {
		return nil, err
	}

	if !common.IsHexAddress(addressStr) {
		return nil, fmt.Errorf("invalid fact address: %v", addressStr)
	}

	addr := common.HexToAddress(addressStr)
	return &addr, nil
}

func parseBigInt(r io.Reader) (res *big.Int, err error) {
	var s string

	if s, err = readLine(r); err != nil {
		return nil, err
	}

	var ok bool
	if res, ok = new(big.Int).SetString(s, 0); !ok {
		return nil, fmt.Errorf("failed to parse number: %v", s)
	}
	return
}
