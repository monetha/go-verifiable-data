# upgrade-passport-logic

Utility tool to upgrade [PassportLogic](../../contracts/code/PassportLogic.sol) contract in [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract.

## Usage

Usage of `./upgrade-passport-logic`:
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
  -version string
    	The version of new passport logic contract (which will be deployed)
  -vmodule string
    	log verbosity pattern
```


## Examples

* Deploying and registering new [PassportLogic](../../contracts/code/PassportLogic.sol) contract of version `0.2` (can be arbitrary text) in simulated environment 
  (for testing) using Ethereum private key stored in file `./owner.key`.
  ```bash
  ./upgrade-passport-logic -ownerkey ./owner.key -version 0.2
  ```
  
* Deploying and registering new [PassportLogic](../../contracts/code/PassportLogic.sol) contract of version `0.2` in 
  [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract (`0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c`) in Ropsten network using Ethereum private 
  key stored in file `./owner.key` (you need to provide private key of [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract owner):
  ```bash
  ./upgrade-passport-logic -ownerkey ./owner.key \
     -registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     -version 0.2 \
     -backendurl https://ropsten.infura.io 
  ```  