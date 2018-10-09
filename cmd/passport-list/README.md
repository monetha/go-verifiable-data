# passport-list

Utility tool for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

## Usage

Usage of `./passport-list`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -factoryaddr string
    	Ethereum address of passport factory contract
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
```

## Examples

* Get example passport from the simulated backend:
  ```bash
  ./passport-list
  ```

* Get all passports created by the passport factory [`0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2#code) in Ropsten network:
  ```bash
  ./passport-list -backendurl https://ropsten.infura.io -factoryaddr 0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
  ```