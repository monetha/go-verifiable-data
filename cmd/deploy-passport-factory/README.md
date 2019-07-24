# deploy-passport-factory

Utility tool to deploy only [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

## Usage

Usage of `./deploy-passport-factory`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -ownerkey string
    	Monetha owner private key filename
  -ownerkeyhex string
    	Monetha owner private key as hex (for testing)
  -registryaddr string
    	Ethereum address of passport logic registry contract
  -v	outputs the binary version
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
  -quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  -quorum_enclave
        Quorum enclave url for private transactions
```


## Examples

* Deploying new [PassportFactory](../../contracts/code/PassportFactory.sol) contract in simulated environment 
  (for testing) using Ethereum private key stored in file `./owner.key`.
  ```bash
  ./deploy-passport-factory -ownerkey ./owner.key
  ```
  
* Deploying new [PassportFactory](../../contracts/code/PassportFactory.sol) contract using existing
  [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract (`0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c`) in Ropsten network using Ethereum private 
  key stored in file `./owner.key`:
  ```bash
  ./deploy-passport-factory -ownerkey ./owner.key \
     -registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     -backendurl https://ropsten.infura.io 
  ```  