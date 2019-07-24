# read-fact

Utility tool to read facts from passport.

## Usage

Usage of `./read-fact`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -factprovideraddr value
    	Ethereum address of fact provider
  -fkey string
    	the key of the fact (max. 32 bytes)
  -ftype string
    	the data type of fact (txdata, string, bytes, address, uint, int, bool, ipfs)
  -ipfsurl string
    	IPFS node address (default "https://ipfs.infura.io:5001")
  -out string
    	save retrieved data to the specified file
  -passportaddr value
    	Ethereum address of passport contract
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
  -quorum_enclave
        Quorum enclave url to decrypt facts, stored using private transactions
```

## Examples

* Retrieve the value of type `txdata` stored under the key `some_key` in simulated backend and write it to the file
  `/dev/stdout` (outputs to the screen):
  ```bash
  ./read-fact -out /dev/stdout \
    -fkey some_key \
    -ftype txdata
  ```
  
* Retrieve the value of type `txdata` stored under the key `monetha.jpg` by the fact provider `0x5b2ae3b3a801469886bb8f5349fc3744caa6348d`
  from passport 
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `./fact_image.jpg`:
  ```bash
  ./read-fact -out ./fact_image.jpg \
    -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    -factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
    -fkey monetha.jpg \
    -ftype txdata \
    -backendurl https://ropsten.infura.io
  ```
  
* Retrieve the value of type `ipfs` stored under the key `Monetha_WP.pdf` by the fact provider `0x5b2ae3b3a801469886bb8f5349fc3744caa6348d`
  from passport 
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `./Monetha_WP.pdf`:
  ```bash
  ./read-fact -out ./Monetha_WP.pdf \
    -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    -factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
    -fkey Monetha_WP.pdf \
    -ftype ipfs \
    -backendurl https://ropsten.infura.io
  ```
