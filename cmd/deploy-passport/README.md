deploy-passport
---------------

Utility tool to deploy passport contract.

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
  [0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2#code) 
  and Ethereum private key stored in file `./owner.key`:
  ```bash
  ./deploy-passport -ownerkey ./owner.key -factoryaddr 0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2 -backendurl https://ropsten.infura.io
  ```