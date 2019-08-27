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

## bootstrap

Command to deploy three contracts at once:

1. [PassportLogic](../../contracts/code/PassportLogic.sol) contract
1. [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract
1. [PassportFactory](../../contracts/code/PassportFactory.sol) contract

After passport factory contract is created, it can be used to deploy [Passport](../../contracts/code/Passport.sol) contracts using
[create](#create) command.

### Usage

Usage of `bootstrap`:
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
  ./passport bootstrap --ownerkey ./owner.key --backendurl https://ropsten.infura.io
  ```

## deploy-factory

Command to deploy only [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

### Usage

Usage of `deploy-factory`:
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
  ./passport deploy-factory --ownerkey ./owner.key \
     --registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     --backendurl https://ropsten.infura.io
  ```

## create

Command to deploy [Passport](../../contracts/code/Passport.sol) contract.

### Usage

Usage of `create`:
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
  ./passport create --ownerkey ./owner.key --factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2 --backendurl https://ropsten.infura.io
  ```

## list

Command for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

### Usage

Usage of `list`:
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
  ./passport list --out ./passports.csv \
    --backendurl https://ropsten.infura.io \
    --factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
  ```
## permission

A command that allows a digital identity owner to allow or deny a data source to write/delete data to/from a digital identity.
By default any data source can write to a digital identity, but a digital identity owner can change permissions that allow only
data sources from the whitelist to write to a digital identity.

## Usage

Usage of `permission`:
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
    ./passport permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --add 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
removes data source `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

    ```bash
    ./passport permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --remove 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows to store the facts only to data sources from the whitelist:

    ```bash
    ./passport permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --enablewhitelist \
      --backendurl https://ropsten.infura.io
    ```

* Owner of a digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows any data source to write the facts:

    ```bash
    ./passport permission --ownerkey pass_owner.key \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --disablewhitelist \
      --backendurl https://ropsten.infura.io
    ```

## upgrade-logic

Command to upgrade [PassportLogic](../../contracts/code/PassportLogic.sol) contract in [PassportLogicRegistry](../../contracts/code/PassportLogicRegistry.sol) contract.

### Usage

Usage of `upgrade-logic`:
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
  ./passport upgrade-logic --ownerkey ./owner.key \
     --registryaddr 0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c \
     --newversion 0.2 \
     --backendurl https://ropsten.infura.io
  ```

## history

Command for reading the history of digital identity changes.

### Usage

Usage of `history`:
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

### Examples

* Read the entire change history for the digital identity [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen):
    ```bash
    ./passport history --out /dev/stdout \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --backendurl https://ropsten.infura.io
    ```

## read-fact

Command to read facts from digital identity.

### Usage

Usage of `read-fact`:
```
  --backendurl string
      backend URL
  --out string
      save retrieved data to the specified file
  --passaddr value
      Ethereum address of passport contract
  --factprovideraddr string
      Ethereum address of data source (fact provider)
  --fkey string
      the key of the fact
  --ftype string
      the data type of fact (txdata, string, bytes, address, uint, int, bool, ipfs, privatedata)
  --ipfsurl string
      IPFS node address (default "https://ipfs.infura.io:5001") (to read ipfs and privatedata facts)
  --ownerkey string
      digital identity owner private key filename (only for privatedata data type)
  --datakey string
      data decryption key file name (only for privatedata data type)
  --quorum_enclave
      Quorum enclave url to decrypt facts, stored using private transactions
  --verbosity int
      log verbosity (0-9) (default 2)
  --vmodule string
      log verbosity pattern
```

### Examples

* Retrieve the value of type `txdata` stored under the key `monetha.jpg` by the data source `0x5b2ae3b3a801469886bb8f5349fc3744caa6348d`
  from digital identity
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `./fact_image.jpg`:
  ```bash
  ./passport read-fact --out ./fact_image.jpg \
    --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    --factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
    --fkey monetha.jpg \
    --ftype txdata \
    --backendurl https://ropsten.infura.io
  ```

* Retrieve the value of type `ipfs` stored under the key `Monetha_WP.pdf` by the data source `0x5b2ae3b3a801469886bb8f5349fc3744caa6348d`
  from digital identity
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `./Monetha_WP.pdf`:
  ```bash
  ./passport read-fact --out ./Monetha_WP.pdf \
    --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    --factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
    --fkey Monetha_WP.pdf \
    --ftype ipfs \
    --backendurl https://ropsten.infura.io
  ```

## read-fact-tx

Command for reading digital identity fact value using transaction hash.

### Usage

Usage of `read-fact-tx`:
```
  --backendurl string
      backend URL
  --out string
      save retrieved data to the specified file
  --passaddr value
      Ethereum address of passport contract
  --ftype string
      the data type of fact (txdata, string, bytes, address, uint, int, bool, ipfs, privatedata)
  --txhash value
      the transaction hash to read history value from
  --ipfsurl string
      IPFS node address (default "https://ipfs.infura.io:5001") (to read ipfs and privatedata facts)
  --verbosity int
      log verbosity (0-9) (default 2)
  --vmodule string
      log verbosity pattern
  --ownerkey string
      digital identity owner private key filename (only for privatedata data type)
  --datakey string
      data decryption key file name (only for privatedata data type)
  --quorum_enclave
      Quorum enclave url to decrypt facts, stored using private transactions
  --verbosity int
      log verbosity (0-9) (default 2)
  --vmodule string
      log verbosity pattern
```

### Examples

* Retrieve the history value of type `txdata` stored in transaction [`0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632`](https://ropsten.etherscan.io/tx/0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632) from digital identity
    [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./history_image.jpg`:
    ```bash
    ./passport read-fact-tx --out history_image.jpg \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --ftype txdata \
      --txhash 0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632 \
      --backendurl https://ropsten.infura.io
    ```

* Retrieve the history value of type `ipfs` stored in transaction [`0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca`](https://ropsten.etherscan.io/tx/0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca) from digital identity
    [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./Monetha_WP.pdf`:
    ```bash
    ./passport read-fact-tx --out Monetha_WP.pdf \
      --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      --ftype ipfs \
      --txhash 0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca \
      --backendurl https://ropsten.infura.io
    ```

## write-fact

Command to write data facts to digital identity.

### Usage

Usage of `write-fact`:
```
  --backendurl string
      backend URL
  --passaddr value
      Ethereum address of passport contract
  --factproviderkey string
      data source (fact provider) private key filename
  --fkey string
      the key of the fact (max 32 bytes)
  --ftype string
      the data type of fact (txdata, string, bytes, address, uint, int, bool, ipfs, privatedata)
  --ipfsurl string
      IPFS node address (default "https://ipfs.infura.io:5001") (to write ipfs and privatedata facts)
  --datakey string
      save data encryption key to the specified file (only for privatedata data type)
  --verbosity int
      log verbosity (0-9) (default 2)
  --vmodule string
      log verbosity pattern
  --quorum_privatefor
      Quorum nodes public keys to make transaction private for, separated by commas
  --quorum_enclave
      Quorum enclave url for private transactions
```

### Gas usage

Cumulative gas usage to store number of character of `a` under the key
`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa` using different data types:

| Number of characters |     `ipfs`, gas used    |     `txdata`, gas used    |  `bytes`, gas used |  `string`, gas used |
|---------------------:|--------------------------:|--------------------------:|-------------------:|-------------------:|
| 10 | 114245 | 70436 | 71079 | 71211 |
| 100 | 114245 | 76598 | 157571 | 157703 |
| 500 | 114245 | 103870 | 425756 | 425888 |
| 1000 | 114245 | 138016 | 781119 | 781251 |
| 5000 | 114245 | 410814 | 3563467 | 3563599 |
| 10000 | 114245 | 751864 | 7036521 | 7036653 |
| 50000 | 114245 | 3483963 | - | - |
| 100000 | 114245 | 6907662 | - | - |
| 110000 | 114245 | 7593621 | - | - |
| 120000 | 114245 | 8279814 | - | - |
| 130000 | 114245 | 8966537 | - | - |

### Examples

* Store image from the file `~/Downloads/monetha.jpg` under the key `monetha.jpg` as `txdata` in digital identity
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064):
  ```bash
  ./passport write-fact --factproviderkey fact_provider.key \
    --fkey monetha.jpg \
    --ftype txdata \
    --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    --backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
  ```

* Store image from the file `~/Downloads/monetha.jpg` under the key `monetha.jpg` as `ipfs` (data will be stored in IPFS,
  only hash will be stored in the Ethereum storage) in digital identity
  [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064):
  ```bash
  ./passport write-fact --factproviderkey fact_provider.key \
    --fkey monetha.jpg \
    --ftype ipfs \
    --passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
    --backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
  ```