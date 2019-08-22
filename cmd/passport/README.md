# Digital identity management tool

This is command line utility which provides various commands on digital identities.

## Usage

`./passport <COMMAND> [COMMAND_ARGUMENTS]`

See commands below for individual usage or access command line help by running:

`./passport --help`

and for specific command:

`./passport <COMMAND> --help`

To output version of tool, run:

`./passport --version`

# Commands

## deploy-bootstrap

Command to deploy three contracts at once:

1. [PassportLogic](../../contracts/code/PassportLogic.sol) contract
1. [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract
1. [PassportFactory](../../contracts/code/PassportFactory.sol) contract

After passport factory contract is created, it can be used to deploy [Passport](../../contracts/code/Passport.sol) contracts using
[deploy-passport](#deploy-passport) command.

### Usage

Usage of `deploy-bootstrap`:
```
  --backendurl string
    	backend URL (simulated backend used if empty)
  --ownerkey string
    	owner private key filename
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
  --quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  --quorum_enclave
        Quorum enclave url for private transactions
```

### Examples

* Deploying all contracts in Ropsten network using Ethereum private key stored in file `./owner.key`:
  ```bash
  ./passport deploy-bootstrap --ownerkey ./owner.key --backendurl https://ropsten.infura.io
  ```

## deploy-passport-factory

Command to deploy only [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

### Usage

Usage of `deploy-passport-factory`:
```
  --backendurl string
    	backend URL
  --ownerkey string
    	Monetha owner private key filename
  --registryaddr string
    	Ethereum address of passport logic registry contract
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
  --quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  --quorum_enclave
        Quorum enclave url for private transactions
```

### Examples

* Deploying new [PassportFactory](../../contracts/code/PassportFactory.sol) contract using existing
  [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract (`0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c`) in Ropsten network using Ethereum private
  key stored in file `./owner.key`:
  ```bash
  ./passport deploy-passport-factory --ownerkey ./owner.key \
     --registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     --backendurl https://ropsten.infura.io
  ```

## deploy-passport

Command to deploy [Passport](../../contracts/code/Passport.sol) contract.

### Usage

Usage of `deploy-passport`:
```
  --backendurl string
    	backend URL
  --factoryaddr string
    	Ethereum address of passport factory contract
  --ownerkey string
    	owner private key filename
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
  --quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  --quorum_enclave
        Quorum enclave url for private transactions
```

### Examples

* Deploying passport contract in Ropsten network using passport factory contract at
  [0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2](https://ropsten.etherscan.io/address/0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2#code)
  and Ethereum private key stored in file `./owner.key`:
  ```bash
  ./passport deploy-passport --ownerkey ./owner.key --factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2 --backendurl https://ropsten.infura.io
  ```

## passport-list

Command for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

### Usage

Usage of `passport-list`:
```
  --backendurl string
    	backend URL
  --factoryaddr string
    	Ethereum address of passport factory contract
  --out string
    	save retrieved passports to the specified file
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
```

### Examples

* Get all passports created by the passport factory [`0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2`](https://ropsten.etherscan.io/address/0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2#code) in Ropsten network
  and write them to the file `./passports.csv`:
  ```bash
  ./passport passport-list --out ./passports.csv \
    --backendurl https://ropsten.infura.io \
    --factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
  ```
