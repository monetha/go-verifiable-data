# deploy-bootstrap

Utility tool to deploy three contracts at once:

1. [PassportLogic](../../contracts/code/PassportLogic.sol) contract
1. [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract
1. [PassportFactory](../../contracts/code/PassportFactory.sol) contract

After passport factory contract is created, it can be used to deploy [Passport](../../contracts/code/Passport.sol) contracts using 
[deploy-passport](../deploy-passport) tool.

## Usage

Usage of `./deploy-bootstrap`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -ownerkey string
    	owner private key filename
  -ownerkeyhex string
    	private key as hex (for testing)
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

* Deploying all contracts in simulated environment (for testing) using Ethereum private key stored in file `./owner.key`.
  ```bash
  ./deploy-bootstrap -ownerkey ./owner.key
  ```

* Deploying all contracts in Ropsten network using Ethereum private key stored in file `./owner.key`:
  ```bash
  ./deploy-bootstrap -ownerkey ./owner.key -backendurl https://ropsten.infura.io
  ```
 