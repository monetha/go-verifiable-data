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
    	contract owner private key filename
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
    	contract owner private key filename
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
## passport-permission

A command that allows a digital identity owner to allow or deny a data source to write/delete data to/from a digital identity.
By default any data source can write to a digital identity, but a digital identity owner can change permissions that allow only
data sources from the whitelist to write to a digital identity.

## Usage

Usage of `passport-permission`:
```
  --add string
    	add data source address to the whitelist
  --remove string
    	remove data source address from the whitelist
  --backendurl string
    	backend URL
  --enablewhitelist
    	enables the use of a whitelist of data sources
  --disablewhitelist
        disables the use of a whitelist of data sources
  --ownerkey string
    	owner private key filename
  --passaddr string
    	Ethereum address of passport contract
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
  --quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  --quorum_enclave
        Quorum enclave url for private transactions
```

## Examples

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
adds data source `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

    ```bash
    ./passport passport-permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --add 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
removes data source `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

    ```bash
    ./passport passport-permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --remove 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows to store the facts only to data sources from the whitelist:

    ```bash
    ./passport passport-permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --enablewhitelist \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows any data source to write the facts:

    ```bash
    ./passport passport-permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --disablewhitelist \
      --backendurl https://ropsten.infura.io
    ```

## upgrade-passport-logic

Command to upgrade [PassportLogic](../../contracts/code/PassportLogic.sol) contract in [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract.

### Usage

Usage of `upgrade-passport-logic`:
```
  --backendurl string
    	backend URL
  --ownerkey string
    	registry contract owner private key filename
  --registryaddr string
    	Ethereum address of passport logic registry contract
  --newversion string
    	The version of new passport logic contract (which will be deployed)
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

* Deploying and registering new [PassportLogic](../../contracts/code/PassportLogic.sol) contract of version `0.2` in
  [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract (`0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c`) in Ropsten network using Ethereum private
  key stored in file `./owner.key` (you need to provide private key of [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract owner):
  ```bash
  ./passport upgrade-passport-logic --ownerkey ./owner.key \
     --registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     --newversion 0.2 \
     --backendurl https://ropsten.infura.io
  ```

## read-history

Command for reading the history of digital identity changes.

## Usage

Usage of `read-history`:
```
  --backendurl string
    	backend URL
  --out string
    	save retrieved data to the specified file
  --passaddr value
    	Ethereum address of passport contract
  --verbosity int
    	log verbosity (0-9) (default 2)
  --vmodule string
    	log verbosity pattern
```

## Examples

* Read the entire change history for the digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen):
    ```bash
    ./passport read-history --out /dev/stdout \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --backendurl https://ropsten.infura.io
    ```