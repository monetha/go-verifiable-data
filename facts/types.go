package facts

// PrivateDataHashes holds IPFS hash of the data and hash of data encryption key
type PrivateDataHashes struct {
	DataIPFSHash string
	DataKeyHash  [32]byte
}
