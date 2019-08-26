package cmdutils

import "github.com/monetha/go-verifiable-data/cmd/internal/flag"

// QuorumPrivateTxICommand - provides commands, needed for Quorum's private tx decrypting
type QuorumPrivateTxICommand struct {
	QuorumEnclave string `long:"quorum_enclave" description:"Quorum enclave url for private transactions"`
}

// QuorumPrivateTxIOCommand - provides commands, needed for Quorum's private tx writing and decrypting
type QuorumPrivateTxIOCommand struct {
	QuorumPrivateTxICommand
	QuorumPrivateFor flag.StringArray `long:"quorum_privatefor" description:"Quorum nodes public keys to make transaction private for, separated by commas"`
}
