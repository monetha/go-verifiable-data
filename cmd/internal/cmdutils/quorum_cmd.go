package cmdutils

import "github.com/monetha/reputation-go-sdk/cmd/internal/flag"

// QuorumPrivateTxIOCommand - provides commands, needed for Quorum's private tx writing and decrypting
type QuorumPrivateTxIOCommand struct {
	QuorumEnclave string `long:"quorum_enclave" description:"Quorum enclave url for private transactions"`
	QuorumPrivateFor flag.StringArray `long:"quorum_privatefor" description:"Quorum nodes public keys to make transaction private for, separated by commas"`
}
