# deploy-passport

Utility tool to deploy [Passport](../../contracts/code/Passport.sol) contract.

## Usage

Usage of `./deploy-passport`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -factoryaddr string
    	Ethereum address of passport factory contract
  -ownerkey string
    	owner private key filename
  -ownerkeyhex string
    	private key as hex (for testing)
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
```

## Examples

* Deploying passport contract in simulated environment (for testing) using Ethereum private key stored in file `./owner.key`.
  ```bash
  ./deploy-passport -ownerkey ./owner.key
  ```

* Deploying passport contract in Ropsten network using passport factory contract at
  [0x5FD962855e9b327262F47594949fd6d742FE2A01](https://ropsten.etherscan.io/address/0x5FD962855e9b327262F47594949fd6d742FE2A01#code) 
  and Ethereum private key stored in file `./owner.key`:
  ```bash
  ./deploy-passport -ownerkey ./owner.key -factoryaddr 0x5FD962855e9b327262F47594949fd6d742FE2A01 -backendurl https://ropsten.infura.io
  ```