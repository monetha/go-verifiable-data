//go:generate go get github.com/alvaroloes/enumer
//go:generate enumer -type=Type

package data

// Type is an enumeration for data type
type Type uint

const (
	// TxData is equal to Bytes data type but uses transaction data to store the data
	TxData Type = iota
	// String stored in Ethereum storage
	String Type = iota
	// Bytes stored in Ethereum storage
	Bytes Type = iota
	// Address stored in Ethereum storage
	Address Type = iota
	// Uint stored in Ethereum storage
	Uint Type = iota
	// Int stored in Ethereum storage
	Int Type = iota
	// Bool stored in Ethereum storage
	Bool Type = iota
	// IPFS hash stored in Ethereum storage, data stored in IPFS
	IPFS Type = iota
)
