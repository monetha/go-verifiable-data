package facts

// PrivateData holds IPFS hash of the data and hash of data encryption key
type PrivateData struct {
	DataIPFSHash string
	DataKeyHash  [32]byte
}
