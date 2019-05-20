# Monetha: Decentralized Reputation Framework [![GoDoc][1]][2] [![Build Status][3]][4] [![Go Report Card][5]][6] [![Coverage Status][7]][8]

[1]: https://godoc.org/github.com/monetha/reputation-go-sdk?status.svg
[2]: https://godoc.org/github.com/monetha/reputation-go-sdk
[3]: https://travis-ci.org/monetha/reputation-go-sdk.svg?branch=master
[4]: https://travis-ci.org/monetha/reputation-go-sdk
[5]: https://goreportcard.com/badge/github.com/monetha/reputation-go-sdk
[6]: https://goreportcard.com/report/github.com/monetha/reputation-go-sdk
[7]: https://codecov.io/gh/monetha/reputation-go-sdk/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/monetha/reputation-go-sdk

# Reputation Layer: go-sdk

* [Building the source](#building-the-source)
    * [Prerequisites](#prerequisites)
    * [Build](#build)
    * [Executables](#executables)
* [Contributing](#contributing)
    * [Making changes](#making-changes)
    * [Contracts update](#contracts-update)
    * [Formatting source code](#formatting-source-code)
* [Bootstrap reputation layer](#bootstrap-reputation-layer)
* [Usage](#usage)
    * [Deploying passport](#deploying-passport)
    * [Passport list](#passport-list)
    * [Writing facts](#writing-facts)
    * [Reading facts](#reading-facts)
    * [Changing passport permissions](#changing-passport-permissions)
    * [Reading facts history](#reading-facts-history)
    * [Passport scanner](#passport-scanner) (sample web application)
    * [Private data](#private-data)        
        * [Writing private data](#writing-private-data)
        * [Reading private data](#reading-private-data)
        * [Reading history of private data](#reading-history-of-private-data)
    * [Private data exchange](#private-data-exchange)
        * [Proposing private data exchange](#proposing-private-data-exchange)
        * [Getting status of private data exchange](#getting-status-of-private-data-exchange)
        * [Accepting private data exchange](#accepting-private-data-exchange)
        * [Reading private data after private data exchange acceptance](#reading-private-data-after-private-data-exchange-acceptance)
        * [Closing private data exchange proposition when timed out](#closing-private-data-exchange-proposition-when-timed-out)        
        * [Closing private data exchange after acceptance](#closing-private-data-exchange-after-acceptance)
        * [Opening dispute after private data exchange acceptance](#opening-dispute-after-private-data-exchange-acceptance)
* [Permissioned blockchains support](#permissioned-blockchains-support)
    * [Quorum](#quorum)

## Building the source

### Prerequisites

1. Make sure you have [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed.
1. Install [Go 1.12](https://golang.org/dl/)
1. Setup `$GOPATH` environment variable as described [here](https://github.com/golang/go/wiki/SettingGOPATH).
1. Clone the repository:
    ```bash
    mkdir -p $GOPATH/src/github.com/monetha
    cd $GOPATH/src/github.com/monetha
    git clone git@github.com:monetha/reputation-go-sdk.git
    cd reputation-go-sdk
    ```

**Note**: You can skip steps 2-3 on Linux and use the official docker image for Go after step 4 to build the project:

```bash
docker run -it --rm \
  -v "$PWD":/go/src/github.com/monetha/reputation-go-sdk \
  -w /go/src/github.com/monetha/reputation-go-sdk \
  golang:1.12 \
  /bin/bash
```

### Build

Install dependencies:

    make dependencies

Once the dependencies are installed, run 

    make cmd

to build the full suite of utilities. After the executable files are built, they can be found in the directory `./bin/`.

### Executables

The reputation-go-sdk project comes with several executables found in the [`cmd`](cmd) directory.

| Command    | Description |
|:----------:|-------------|
| [`deploy-bootstrap`](cmd/deploy-bootstrap) | Utility tool to deploy three contracts at once: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol). |
| [`deploy-passport`](cmd/deploy-passport) | Utility tool to deploy [Passport](contracts/code/Passport.sol) contracts using already deployed [PassportFactory](contracts/code/PassportFactory.sol). |
| [`upgrade-passport-logic`](cmd/upgrade-passport-logic) | Utility tool to upgrade [PassportLogic](contracts/code/PassportLogic.sol) contract. |
| [`write-fact`](cmd/write-fact) | Utility tool to write facts to passport. |
| [`read-fact`](cmd/read-fact) | Utility tool to read facts from passport. |
| [`passport-list`](cmd/passport-list) | Utility tool for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract. |
| [`passport-permission`](cmd/passport-permission) | Utility tool that allows a passport holder to allow or deny a fact provider to write/delete facts to/from a passport. By default any fact provider can write to a passport, but a passport holder can change permissions that allow only fact providers from the whitelist to write to a passport. |
| [`read-history`](cmd/read-history) | Utility tool for reading the history of passport changes. |
| [`passport-scanner`](cmd/passport-scanner) | Web application (WebAssembly module) to get the list of deployed passports and the history of passport changes in a web browser. |

## Contributing

### Making changes

Make your changes, then ensure that the tests and the linters pass:

    make lint
    make test

To get the test coverage report for all packages, run the command:

    make cover

Тest coverage results (`cover.out`, `cover.html`) will be put in `./.cover` directory.

### Contracts update

After Ethereum contracts code is updated and artifacts are created:
1. Copy all artifacts to [`contracts/code`](contracts/code) folder.
1. Run `go generate` command in [`contracts`](contracts) folder to convert Ethereum contracts into Go package.
1. Commit new/updated files.

### Formatting source code

`make fmt` command automatically formats Go source code of the entire project.

## Bootstrap reputation layer

Monetha has already deployed this set of auxiliary reputation layer contracts on Ropsten test network and Mainnet network. 

The contract addresses deployed on Ropsten:

| Contract      | Address                                      |
|---------------|----------------------------------------------|
| `PassportLogic` | [`0x74C22a5d68E4727029FD906aD73D5F39D9130905`](https://ropsten.etherscan.io/address/0x74C22a5d68E4727029FD906aD73D5F39D9130905) |
| `PassportLogicRegistry`  | [`0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c`](https://ropsten.etherscan.io/address/0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c) |
| `PassportFactory` | [`0x5FD962855e9b327262F47594949fd6d742FE2A01`](https://ropsten.etherscan.io/address/0x5FD962855e9b327262F47594949fd6d742FE2A01) |

The contract addresses deployed on Mainnet:

| Contract      | Address                                      |
|---------------|----------------------------------------------|
| `PassportLogic` | [`0x76E2fe5C37c47Fe09DCFa55Bec9Fd34318922F27`](https://etherscan.io/address/0x76E2fe5C37c47Fe09DCFa55Bec9Fd34318922F27) |
| `PassportLogicRegistry`  | [`0x41c32A8387ff178659ED9B04190613623F545657`](https://etherscan.io/address/0x41c32A8387ff178659ED9B04190613623F545657) |
| `PassportFactory` | [`0xdbf780f836D8a22b56AcF9Fd266171fAFf31F521`](https://etherscan.io/address/0xdbf780f836D8a22b56AcF9Fd266171fAFf31F521) |

Consider the process of deploying your own set of auxiliary repoutation layer contracts to experiment with our implementation. If you are going to deploy your contracts, then you will have to support them yourself.

This means that if the reputation layer logic of the passport is updated by Monetha developers, you'll need to deploy a new `PassportLogic` contract, register it 
in an existing `PassportLogicRegistry` contract (by calling `addPassportLogic` method) and finally make it active (by calling `setCurrentPassportLogic`).

If you use a set of Monetha deployed reputation layer contracts, then the reputation passport logic is always up-to-date with latest fixes and features.

Prepare in advance the address that will be the owner of the deployed contracts. Make sure that it has enough funds to deploy contracts (1 ETH should be enough).
Store the private key of this address in the file named `./owner.key`.

To deploy all contracts in Rospten network using Ethereum private key stored in file `./owner.key`, run the command:

```bash
./bin/deploy-bootstrap -ownerkey ./owner.key \
  -backendurl https://ropsten.infura.io
```
    
After running the command, you should see something like the following output:

```
WARN [03-19|10:29:04.273] Loaded configuration                     owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 backend_url=https://ropsten.infura.io
WARN [03-19|10:29:05.038] Getting balance                          address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [03-19|10:29:05.172] Deploying PassportLogic                  owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [03-19|10:29:05.759] Waiting for transaction                  hash=0xaa4149349dc9856289cacdb9158d6d215124f7a9b3d182c3b31429004a908bc8
WARN [03-19|10:29:18.216] Transaction successfully mined           tx_hash=0xaa4149349dc9856289cacdb9158d6d215124f7a9b3d182c3b31429004a908bc8 cumulative_gas_used=2412986
WARN [03-19|10:29:18.216] PassportLogic deployed                   contract_address=0xEf95422e66761A5a468FE72c1fD3C946884d5E50
WARN [03-19|10:29:18.216] Deploying PassportLogicRegistry          owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 impl_version=0.1 impl_address=0xEf95422e66761A5a468FE72c1fD3C946884d5E50
WARN [03-19|10:29:18.704] Waiting for transaction                  hash=0x1b0174592e60e8e7f8619b6e1d9bceb498530feb405c7f8f6715c26d68445ebb
WARN [03-19|10:29:39.585] Transaction successfully mined           tx_hash=0x1b0174592e60e8e7f8619b6e1d9bceb498530feb405c7f8f6715c26d68445ebb cumulative_gas_used=2157734
WARN [03-19|10:29:39.585] PassportLogicRegistry deployed           contract_address=0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c
WARN [03-19|10:29:39.585] Deploying PassportFactory                owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 registry=0x11C96d40244d37ad3Bb788c15F6376cEfA28CF7c
WARN [03-19|10:29:40.275] Waiting for transaction                  hash=0x984b1df10640a5fe3005aea593b1cbb7af59682469564063a490820c8a8cb6f0
WARN [03-19|10:29:48.602] Transaction successfully mined           tx_hash=0x984b1df10640a5fe3005aea593b1cbb7af59682469564063a490820c8a8cb6f0 cumulative_gas_used=1446005
WARN [03-19|10:29:48.602] PassportFactory deployed                 contract_address=0x5FD962855e9b327262F47594949fd6d742FE2A01
WARN [03-19|10:29:48.602] Done.
```

In the output you can find the addresses of all the deployed contracts.


## Usage

In order to create a passport and start using it, you need to use auxiliary reputation layer contracts: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol).

### Deploying passport

Before creating a passport for a specific Ethereum address, store the private key of this Ethereum address in the file `pass_owner.key`.
Make sure that the passport owner has enough money to create a passport contract. Usually passport contract deployment takes `425478` gas.

To create a passport contract you need to know address of the `PassportFactory` contract. Let's try to create a passport in Ropsten
using the `PassportFactory` contract deployed by Monetha ([`0x5FD962855e9b327262F47594949fd6d742FE2A01`](https://ropsten.etherscan.io/address/0x5FD962855e9b327262F47594949fd6d742FE2A01)):

```bash
./bin/deploy-passport -ownerkey ./pass_owner.key \
  -factoryaddr 0x5FD962855e9b327262F47594949fd6d742FE2A01 \
  -backendurl https://ropsten.infura.io
```

Below you can see the output of the command to create a passport for the address `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d`.

```
WARN [03-19|10:41:59.024] Loaded configuration                     owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 backend_url=https://ropsten.infura.io factory=0x5FD962855e9b327262F47594949fd6d742FE2A01
WARN [03-19|10:41:59.798] Getting balance                          address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [03-19|10:41:59.933] Initializing PassportFactory contract    factory=0x5FD962855e9b327262F47594949fd6d742FE2A01
WARN [03-19|10:41:59.940] Deploying Passport contract              owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [03-19|10:42:00.692] Waiting for transaction                  hash=0x33e7d5dc34f5e8597859c319c34ef4f613238defbadcc2fda3ae65f508b45884
WARN [03-19|10:42:08.980] Transaction successfully mined           tx_hash=0x33e7d5dc34f5e8597859c319c34ef4f613238defbadcc2fda3ae65f508b45884 cumulative_gas_used=6130011
WARN [03-19|10:42:08.987] Passport deployed                        contract_address=0x1C3A76a9A27470657BcBE7BfB47820457E4DB682
WARN [03-19|10:42:08.987] Initializing Passport contract           passport=0x5FD962855e9b327262F47594949fd6d742FE2A01
WARN [03-19|10:42:08.987] Claiming ownership                       owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [03-19|10:42:09.550] Waiting for transaction                  hash=0xae51f2769db4bb7d5c8c651cfa78b8048f5fcd6bc949ac3d6b220aa7c2d5255e
WARN [03-19|10:42:22.061] Transaction successfully mined           tx_hash=0xae51f2769db4bb7d5c8c651cfa78b8048f5fcd6bc949ac3d6b220aa7c2d5255e cumulative_gas_used=2220151
WARN [03-19|10:42:22.061] Done.
```

As you can see in the line `Passport deployed`, a passport contract was created at address [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x1C3A76a9A27470657BcBE7BfB47820457E4DB682).

### Passport list

The passport factory allows you to get a list of all the passports that have been created.

Let's try to get a list of all passports using the address of `PassportFactory` contract deployed by Monetha ([`0x5FD962855e9b327262F47594949fd6d742FE2A01`](https://ropsten.etherscan.io/address/0x5FD962855e9b327262F47594949fd6d742FE2A01))
in Ropsten network:

```bash
./bin/passport-list -out /dev/stdout \
  -backendurl https://ropsten.infura.io \
  -factoryaddr 0x5FD962855e9b327262F47594949fd6d742FE2A01
```

You should see something like this:

```
WARN [03-19|10:44:48.830] Loaded configuration                     factory_provider=0x5FD962855e9b327262F47594949fd6d742FE2A01 backend_url=https://ropsten.infura.io
WARN [03-19|10:44:49.579] Initialising passport factory contract   passport_factory=0x5FD962855e9b327262F47594949fd6d742FE2A01
WARN [03-19|10:44:49.579] FilterPassportCreated                    start_block=0
WARN [03-19|10:44:49.841] Writing collected passports to file
passport_address,first_owner,block_number,tx_hash
0x1C3A76a9A27470657BcBE7BfB47820457E4DB682,0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55,5233845,0x33e7d5dc34f5e8597859c319c34ef4f613238defbadcc2fda3ae65f508b45884
WARN [03-19|10:44:49.841] Done.
```

The output can be saved to a file and converted to the table. Currently one passport is deployed:

|passport_address|first_owner|block_number|tx_hash|
|----------------|-----------|------------|-------|
|0x1C3A76a9A27470657BcBE7BfB47820457E4DB682|0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55|5233845|0x33e7d5dc34f5e8597859c319c34ef4f613238defbadcc2fda3ae65f508b45884|

The block number and transaction hash indicate the transaction in which the passport was created.

All passports use the same passport logic contract. Once a new passport logic is added to the passport logic registry and is
activated, it will be immediately used by all passports created by this factory. How cool is that!

### Writing facts

After the passport is created, any fact provider can start writing data to the passport.

Before we start writing facts to a passport, let's store the private key of the fact provider to the file `fact_provider.key`.
Make sure that the fact provider has enough funds to write the facts. Check [gas usage table](cmd/write-fact#gas-usage) to estimate the required amount of funds.

You can write up to 100KB of data in passport under one key when `txdata` data type is used. Supported data types that 
can be written to the passport: `string`, `bytes`, `address`, `uint`, `int`, `bool`, `txdata`. All types except `txdata` 
use Ethereum storage to store the data. `txdata` uses Ethereum storage only to save the block number, the data itself 
remains in the transaction input data and can be read later using the SDK. Therefore, if you need to save a large amount 
of data, it is better to use `txdata` type of data. The disadvantage of the `txdata` type of data is the data can only be read 
using the SDK, within the contracts this data is not available.

Let's try to store image from the file `~/Downloads/monetha.jpg` under the key `monetha.jpg` as `txdata` in passport
`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`:

```bash
./bin/write-fact -ownerkey fact_provider.key \
  -fkey monetha.jpg \
  -ftype txdata \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
```

Similarly, you can store the same image from the file `~/Downloads/monetha.jpg` under the key `monetha.jpg` as `ipfs` in the same passport.
Keep in mind, the data will be stored in IPFS, only IPFS hash will be stored in the Ethereum storage:

```bash
./bin/write-fact -ownerkey fact_provider.key \
  -fkey monetha.jpg \
  -ftype ipfs \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
```

### Reading facts

After the fact provider has written the public data to the passport, the data can be read by anyone.
To read the data you need to know: the address of the passport, the address of the fact provider who stored the data, 
the key under which the data was stored and the type of data.

Let's try to retrieve image from passport `0x1C3A76a9A27470657BcBE7BfB47820457E4DB682` that was stored by the fact provider
`0x5b2ae3b3a801469886bb8f5349fc3744caa6348d` under the key `monetha.jpg` as `txdata` data type and write it to the file 
`./fact_image.jpg`:

```bash
./bin/read-fact -out ./fact_image.jpg \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
  -fkey monetha.jpg \
  -ftype txdata \
  -backendurl https://ropsten.infura.io
```

After the data has been read from the Ethereum blockchain and written to the file `./fact_image.jpg`, try to open the image.

To get the same file that was previously saved in IPFS, just change the parameter `-ftype` to `ipfs`:

```bash
./bin/read-fact -out ./ipfs_fact_image.jpg \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
  -fkey monetha.jpg \
  -ftype ipfs \
  -backendurl https://ropsten.infura.io
```

Тhe data will be written to the file `./ipfs_fact_image.jpg`.

### Changing passport permissions

By default any fact provider can write to a passport, but a passport owner can change permissions that allow only
fact providers from the whitelist to write to a passport. To do this, the passport owner must add the authorized fact providers 
to the whitelist and then allow to store the facts only to fact providers from the whitelist.

Consider an example of how owner of a passport `0x1C3A76a9A27470657BcBE7BfB47820457E4DB682` adds fact provider 
`0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

```bash
./bin/passport-permission -ownerkey pass_owner.key \
  -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -add 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
  -backendurl https://ropsten.infura.io
```

Please note that the passport owner’s private key is stored in the file `pass_owner.key`.

After executing the command, any fact provider is still allowed to store the facts in the passport. Let's fix it!

Owner of a passport `0x1C3A76a9A27470657BcBE7BfB47820457E4DB682` may allow to store the facts only to fact providers 
from the whitelist by running the command:

```bash
./bin/passport-permission -ownerkey pass_owner.key \
  -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -onlywhitelist true \
  -backendurl https://ropsten.infura.io
```

More examples can be found [here](cmd/passport-permission#examples).

### Reading facts history

The SDK allows you to see the history of absolutely all changes of facts in the passport.

Let's try to retrieve the entire change history for the passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen, but you can change this to the file name to write to the file):

```bash
./bin/read-history -out /dev/stdout \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -backendurl https://ropsten.infura.io
```

After running the command you should see something like this:

```
WARN [03-19|11:00:42.836] Loaded configuration                     backend_url=https://ropsten.infura.io passport=0x1C3A76a9A27470657BcBE7BfB47820457E4DB682
fact_provider,key,data_type,change_type,block_number,tx_hash
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,TxData,Updated,5233914,0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,IPFS,Updated,5233917,0xf069012520c55d293595654805f3f2b1ff032c1395ddd37cd1366fc1ac67114e
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,Monetha_WP.pdf,IPFS,Updated,5233930,0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca
WARN [03-19|11:00:44.497] Done.
```

The CSV output can be saved to a file and converted to the table:

| fact_provider | key | data_type | change_type | block_number | tx_hash |
|---------------|-----|-----------|-------------|--------------|---------|
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | TxData | Updated | 5233914 | 0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632 |
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | IPFS | Updated | 5233917 | 0xf069012520c55d293595654805f3f2b1ff032c1395ddd37cd1366fc1ac67114e |
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | Monetha_WP.pdf | IPFS | Updated | 5233930 | 0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca |

As we can see, there were two fact updates of type `TxData` (under the same key `monetha.jpg`) by the same data provider `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d`, 
and one update of type `IPFS` by the same data provider.
The `block_number` and `tx_hash` columns allow us to understand in which block and in which transaction the changes were made.
The `change_type` column may contain either `Updated` or `Deleted` values. Even if the value of a fact has been deleted, we can read its value as it was before the deletion.

Let's read what the value of the fact was during the first update. To do this, we need to specify the transaction hash `0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632` and the type of data `txdata`:

```bash
./bin/read-history -out monetha1.jpg \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -ftype txdata \
  -txhash 0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632 \
  -backendurl https://ropsten.infura.io
```

Similarly, we can read what fact value was written in the second transaction `0xf069012520c55d293595654805f3f2b1ff032c1395ddd37cd1366fc1ac67114e`:

```bash
./bin/read-history -out monetha2.jpg \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -ftype ipfs \
  -txhash 0xf069012520c55d293595654805f3f2b1ff032c1395ddd37cd1366fc1ac67114e \
  -backendurl https://ropsten.infura.io
```

Now you can compare pictures `monetha1.jpg` and `monetha2.jpg` to see what changes have been made.

To read fact value of type `TxData` in the third transaction parameter `-ftype` should be changed to `ipfs` and `-txhash` 
to `0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca`:

```bash
./bin/read-history -out Monetha_WP.pdf \
  -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
  -ftype ipfs \
  -txhash 0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca \
  -backendurl https://ropsten.infura.io
```

Тhe value will be written to the file `Monetha_WP.pdf`.

### Passport scanner

After the [go-ethereum](https://github.com/ethereum/go-ethereum) community recently accepted [our changes](https://github.com/ethereum/go-ethereum/pull/17709), 
it became possible to compile Go SDK of reputation layer into WebAssembly module and run it in a browser. 
We have prepared a sample web application that has the same functionality as [`passport-list`](cmd/passport-list) and 
[`read-history`](cmd/read-history) utilities provide, i.e. it allows you to get the list of deployed passports and 
the history of passport changes directly in your web browser.

To play with the web application, run the command

```bash
./bin/passport-scanner
```

and open [http://localhost:8080](http://localhost:8080) in your browser. More details can be found [here](cmd/passport-scanner).

The latest version of passport scanner is also uploaded to IPFS: https://ipfs.io/ipfs/QmNyHrzkD5RpxmxJyJKg9QraUYtGJ48KhskWjnGPfudhoy/

Happy scanning!

### Private data

Private data is stored in encrypted form in IPFS, only the IPFS hash and hash of data encryption key are saved in the 
blockchain.

Reading/writing private data is as simple as reading/writing public data. The only difference is that only the person 
who is the passport owner at the time of writing private data can read the 
private data. The private data provider can read private data only if it has saved the data encryption key.
The passport owner does not need to know the data encryption key, as he can decrypt all private data using his Ethereum 
private key.

#### Writing private data

In order for the fact provider to write private data, it needs to specify the private data type: `-ftype privatedata`.

If the fact provider specifies a `-datakeyfile` parameter, the encryption key will be saved to a file, which will allow 
the fact provider to read the private data later (for verification purposes, for example).

Let's try to store text `this is a very secret message` under the key `secret_message` as `privatedata` in passport 
`0x4026a67a2C4746b94F168bd4d082708f78d7b29f`, and also save the data encryption key in the file `data_enc.key`:

```bash
echo -n "this is a very secret message" | \
  ./bin/write-fact -ownerkey fact_provider.key \
  -fkey secret_message \
  -ftype privatedata \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -datakeyfile data_enc.key \
  -backendurl https://ropsten.infura.io
```

As a result, you can see something like the following output:

```
WARN [05-13|16:51:53.823] Loaded configuration                     fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E backend_url=https://ropsten.infura.io passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f
WARN [05-13|16:51:54.820] Filtering OwnershipTransferred           newOwner=0xD101709569D2dEc41f88d874Badd9c7CF1106AF7
WARN [05-13|16:51:54.975] Getting transaction by hash              tx_hash=0x40768efcf3e6254216bed543433eaeac586fe0ed25b9de04d22e7677cfc980f1
WARN [05-13|16:51:55.115] Writing ephemeral public key to IPFS...
WARN [05-13|16:52:05.050] Ephemeral public key added to IPFS       hash=QmPDZpSfsbU1DKquxhZxEZ6JmWZ61RSa1UiHaqZHDdqMZV size=73
WARN [05-13|16:52:05.050] Writing encrypted message to IPFS...
WARN [05-13|16:52:16.885] Encrypted message added to IPFS          hash=QmfJuqBT7Kqd8Fb4omSDZWQMTpLkj2Gw81uYJm7wDz3Cpt size=53
WARN [05-13|16:52:16.885] Writing message HMAC to IPFS...
WARN [05-13|16:52:17.664] Message HMAC added to IPFS               hash=QmcqcwppRyKc8yQWrkECuixsBgpwyjiirfYUP4Y3TVKYgf size=40
WARN [05-13|16:52:17.664] Creating directory in IPFS...
WARN [05-13|16:52:29.507] Directory created in IPFS                hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm
WARN [05-13|16:52:29.507] Writing private data hashes to Ethereum  passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" ipfs_hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm data_key_hash=0x07bbab23c88ec28c3ddf46ebb8de1a24a278f1c27789cdb0b4fdfd7c5773f2ab
WARN [05-13|16:52:29.652] Writing IPFS private data hashes to passport fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
WARN [05-13|16:52:30.323] Waiting for transaction                  hash=0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6
WARN [05-13|16:52:34.496] Transaction successfully mined           tx_hash=0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6 gas_used=136031
WARN [05-13|16:52:34.496] Writing data encryption key to file      file_name=data_enc.key
WARN [05-13|16:52:34.497] Done.
```

From the output, you can see that private data was saved in IPFS: [QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm](https://ipfs.infura.io/ipfs/QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm)

However, despite the fact that the private data is stored publicly, only the passport owner or the fact provider (only if he's saved the data encryption key) can decrypt it.

#### Reading private data

After the fact provider has written the public data to the passport, the data can be read either by passport owner or by 
fact provider (only if he's saved the data encryption key). To read private data, the following data should be provided:
* the address of the passport (`-passportaddr` parameter)
* the address of the fact provider who stored the private data (`-factprovideraddr` parameter)
* the key under which the private data was stored (`-fkey` parameter)
* if the data is read by the fact provider, he need to specify data encryption key (`-datakeyfile` parameter)
* if the data is read by the owner of the passport, he need to specify his private key (`-ownerkey` parameter)

Let's try to retrieve private data as fact provider using data encryption key stored in file `data_enc.key` from passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` that was stored by the fact provider `0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E` under the key `secret_message` as `privatedata` data type and write it to the standard output:

```bash
./bin/read-fact -out /dev/stdout \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -factprovideraddr 0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E \
  -fkey secret_message \
  -ftype privatedata \
  -datakeyfile data_enc.key \
  -backendurl https://ropsten.infura.io
```

You should see something like the following output:

```
WARN [05-13|17:00:27.686] Loaded configuration                     fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" backend_url=https://ropsten.infura.io passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f
WARN [05-13|17:00:28.549] Reading private data hashes from Ethereum passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
WARN [05-13|17:00:28.549] Getting IPFS private data hashes         fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
WARN [05-13|17:00:28.681] Reading encrypted message from IPFS      hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=encrypted_message
WARN [05-13|17:00:29.230] Reading message HMAC from IPFS           hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=hmac
WARN [05-13|17:00:29.816] Writing data to file
this is a very secret message
```

Now let's try to read the same private data as passport owner. Instead of parameter `-datakeyfile data_enc.key`, parameter `-ownerkey pass_owner.key` is provided:

```bash
./bin/read-fact -out /dev/stdout \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -factprovideraddr 0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E \
  -fkey secret_message \
  -ftype privatedata \
  -ownerkey pass_owner.key \
  -backendurl https://ropsten.infura.io
```

After running the command passport owner should see something like this:

```
WARN [05-13|17:01:32.755] Loaded configuration                     fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" backend_url=https://ropsten.infura.io passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f
WARN [05-13|17:01:33.632] Reading private data hashes from Ethereum passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
WARN [05-13|17:01:33.632] Getting IPFS private data hashes         fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
WARN [05-13|17:01:33.771] Reading ephemeral public key from IPFS   hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=public_key
WARN [05-13|17:01:34.342] Reading encrypted message from IPFS      hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=encrypted_message
WARN [05-13|17:01:34.835] Reading message HMAC from IPFS           hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=hmac
WARN [05-13|17:01:35.314] Writing data to file
this is a very secret message
```

As you can see from the output, the passport owner additionally reads the ephemeral public key from IPFS, which allows 
him to recover the data encryption key using his Ethereum private key.

#### Reading history of private data

Reading the history of private data works for private data in the same way as for public data. You only need to additionally 
specify the data type `-ftype privatedata`, and specify either a data encryption key (`-datakeyfile` parameter) or a 
passport owner’s Ethereum private key (`-ownerkey` parameter).

Let's try to retrieve the entire change history for the passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` in Ropsten blockchain:

```bash
./bin/read-history -out /dev/stdout \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -backendurl https://ropsten.infura.io
```

The output (converted to the table form):

| fact_provider | key | data_type | change_type | block_number | tx_hash |
|---------------|-----|-----------|-------------|--------------|---------|
| 0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E | secret_message | PrivateData | Updated | 5590046 | 0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6 |
| 0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E | some_text | TxData | Updated | 5590236 | 0xc792c69fd59a4146c06ee524262ab929b2af4a895a04df56d7cd1b0bb20af28b |

From the output, you can see that private data was saved in transaction `0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6`.

Let's try to read the private data from this transaction using data encryption key stored in file `data_enc.key`:

```bash
./bin/read-history -out /dev/stdout \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -txhash 0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6 \
  -ftype privatedata \
  -datakeyfile data_enc.key \
  -backendurl https://ropsten.infura.io
```

The output:

```
WARN [05-13|17:05:03.014] Loaded configuration                     backend_url=https://ropsten.infura.io passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f
WARN [05-13|17:05:03.766] Reading data hashes from Ethereum transaction passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f tx_hash=52dcfb…02dea6
WARN [05-13|17:05:03.766] Getting transaction by hash              tx_hash=0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6
WARN [05-13|17:05:03.918] Reading encrypted message from IPFS      hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=encrypted_message
WARN [05-13|17:05:04.464] Reading message HMAC from IPFS           hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=hmac
this is a very secret message
```

Keep in mind that every time a fact provider writes private data, a random data encryption key is used, so you need to 
specify the encryption key that was used in that particular transaction. The passport owner does not need to know all 
the data encryption keys, because he can restore them with his Ethereum private key.

Reading private data from the same transaction by the passport owner looks very similar (instead of `-datakeyfile` 
parameter `-ownerkey` parameter is specified):

```bash
./bin/read-history -out /dev/stdout \
  -passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  -txhash 0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6 \
  -ftype privatedata \
  -ownerkey ./pass_owner.key \
  -backendurl https://ropsten.infura.io
```

The output:

```
WARN [05-13|17:04:00.042] Loaded configuration                     backend_url=https://ropsten.infura.io passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f
WARN [05-13|17:04:00.798] Reading data hashes from Ethereum transaction passport=0x4026a67a2C4746b94F168bd4d082708f78d7b29f tx_hash=52dcfb…02dea6
WARN [05-13|17:04:00.798] Getting transaction by hash              tx_hash=0x52dcfb7591be53cac31bd81fc2c297eb634b79607328969c7375fb481b02dea6
WARN [05-13|17:04:00.957] Reading ephemeral public key from IPFS   hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=public_key
WARN [05-13|17:04:01.500] Reading encrypted message from IPFS      hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=encrypted_message
WARN [05-13|17:04:02.038] Reading message HMAC from IPFS           hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=hmac
this is a very secret message
```

### Private data exchange

Private data exchange engine enables participants to exchange private data via Passports in a secure manner. Anyone can 
request private data from the passport of user. This is achieved by running an interactive protocol between the passport 
owner and the data requester.

How it works:

1. The data requester initiates retrieval of private data from a passport by calling `propose` command. When executing this 
   command, the data requester specifies which fact provider data he wants to read, encrypts exchange key with the passport 
   owner's public key and transfers to the passport the funds that he is willing to pay for the private data.
  
1. The passport owner receives an event from the Ethereum blockchain or directly from the data requester for the data 
   exchange proposition. If he is satisfied with the proposal, he executes the `accept` command. When executing this command, 
   the passport owner encrypts the data encryption key with the exchange key of data requester and 
   transfers the same amount of funds as the data requester to he passport as a guarantee of the validity of the data encryption key.
   
   The data requester has 24 hours to accept private data exchange. 24 hours after the exchange proposition, the data 
   requester can close the proposition and return staked funds back by calling `timeout` command.
   
1. The data requester receives an event from the Ethereum blockchain or directly from the passport owner about accepted 
   private data exchange. It decrypts the data access key using exchange key and reads private data using `read` command. 
   After that `finish` command is called, which returns all staked funds to the passport owner.
   
   During the first 24 hours, the `finish` command can only be called by the data requester, after 24 hours anyone can call this command.

1. If it is not possible to decrypt the data, the data requester calls the `dispute` command, revealing the exchange key.
   The Ethereum contract code identifies the cheater and transfers all staked funds to the party who behaved honestly.
   The data requester has 24 hours to open a dispute, otherwise the exchange is considered valid and the passport owner
   can get all staked funds.

This is how it looks in the state diagram:
   
![PlantUML model](http://www.plantuml.com/plantuml/png/jPF1JWCX48RlFCKSTqtRW_7KWwbH4prfZ3VZWSBiGheB28DjtzujbLGQgscgUmAopFzz0ym2SK-nxvZI4W5xHskG68JNZhGrZBsSlS9uV0cFtZeRKC8Kt7POrSnOGl2wLGJMGAVDWWdUTIXXlfw2vCJ1url4GEXPEPqo6CEGli00jyzt3D_HK5hCIHMkXEAcnNkv6gLYJtdp21mFmLbF3qk3lcPe96nW6Ckx4_IL4EWeGVCq_9KvrmMxASoAwM7c7FGNpDVTPvj9zsZZW0oy8VHmVg4c9tUyHGfR1RbHW3aNYvr72Yyjld9covApqKO7TUHjW4f6hqqxM86Qr0nsd_N0pTeQX2g9vr-AipXiyzswRVRYJrIMEhX8MDMGBKuy6wYM2WsKYY0KSa9P7-dwuoNEKNlvEUfVspeitwJExJ-K48N049hOZROavVkO3SFOTny0)

At any time, the `status` command can be used to get detailed information about the private data exchange.

At all steps of the interactive protocol, the utility `privatedata-exchange` is used.

#### Proposing private data exchange

To initiate the exchange of private data, the data requester must specify the following parameters:

* the passport address (`--passportaddr` parameter)
* the address of the data provider who stored the private data (`--factprovideraddr` parameter)
* key under which private data was stored (`--fkey` parameter)
* the name of the file with the Ethereum private key of data requester (`--requesterkey` parameter)
* the amount of funds (in wei) that the requester is willing to pay for private data (`--stake` parameter)
* the name of the file where the exchange key will be saved (`--exchangekey` parameter), used later both for 
  accessing private data and for resolving a possible dispute

In the example below, the data requester attempts to initiate the retrieval of private data that was stored by the fact provider
`0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E` in passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` under the key `secret_message` 
by staking `10000000000000000 wei` (which is equal to `0.01 ETH`).

```bash
./bin/privatedata-exchange propose \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --factprovideraddr 0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E \
  --fkey secret_message \
  --requesterkey ./data_requester.key \
  --stake 10000000000000000 \
  --exchangekey ./exchange.key \
  --backendurl https://ropsten.infura.io
```

As a result of the command, you can see that the private data exchange proposition was created under index `1` (the index 
is simply the data exchange identifier to refer it in all subsequent commands), and the exchange key was written to file `exchange.key`:

```
WARN [05-16|13:56:31.944] Filtering OwnershipTransferred           newOwner=0xD101709569D2dEc41f88d874Badd9c7CF1106AF7
WARN [05-16|13:56:32.091] Getting transaction by hash              tx_hash=0x40768efcf3e6254216bed543433eaeac586fe0ed25b9de04d22e7677cfc980f1
WARN [05-16|13:56:32.223] Proposing private data exchange          fact_provider=0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E fact_key="[115 101 99 114 101 116 95 109 101 115 115 97 103 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" encrypted_key="[4 206 42 209 27 210 127 104 62 239 227 22 238 102 255 172 187 19 208 6 48 131 70 132 255 136 108 110 176 192 115 160 205 184 242 144 94 240 142 123 21 166 118 215 100 68 146 194 23 163 41 128 240 212 150 188 107 50 216 9 129 110 136 112 25]" key_hash=e86ec7…d8e4b9
WARN [05-16|13:56:32.980] Waiting for transaction                  hash=0xd353c8d21f44a3f17ebe7782bedf03a3b6d4456721f24192bb9813e916977526
WARN [05-16|13:56:37.123] Transaction successfully mined           tx_hash=0xd353c8d21f44a3f17ebe7782bedf03a3b6d4456721f24192bb9813e916977526 gas_used=398976
WARN [05-16|13:56:37.123] PrivateDataExchangeProposed              exchange_index=1 data_requester=0xd2Bb3Aa3F2c0bdA6D8020f3228EabD4A89d8B951 passport_owner=0xD101709569D2dEc41f88d874Badd9c7CF1106AF7
WARN [05-16|13:56:37.123] Private data exchange proposed           exchange_index=1
WARN [05-16|13:56:37.123] Writing exchange key to file             file_name=./exchange.key
```

#### Getting status of private data exchange

`status` command allows to get more detailed information about the private data exchange. To get the information the following parameters should be specified:

* the passport address (`--passportaddr` parameter)
* the private data exchange index (`--exchidx` parameter)

Let's try to get this information about private data exchange referred by the index `1` from passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f`:

```bash
./bin/privatedata-exchange status \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 1 \
  --backendurl https://ropsten.infura.io
```

Immediately after creating the private data exchange proposition, you can see that it's in `Proposed` state, data 
requester staked `0.01 ETH`, and the passport owner has one day left to accept it:

```
Private data exchange:               1 (Proposed, expires in 1 day)
Data requester address:              0xd2Bb3Aa3F2c0bdA6D8020f3228EabD4A89d8B951
Data requester staked:               0.01 ETH
Passport owner address:              0xD101709569D2dEc41f88d874Badd9c7CF1106AF7
Passport owner staked:               0 ETH
Private data fact provider:          0xd8CD4f4640D9Df7ae39aDdF08AE2c6871FcFf77E
Private data fact key:               secret_message
Private data IPFS hash:              QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm
```

#### Accepting private data exchange

To accept the private data exchange after proposition, passport owner should execute `accept` command providing the following parameters:

* the passport address (`--passportaddr` parameter)
* the private data exchange index (`--exchidx` parameter)
* the passport owner's Ethereum private key (`--ownerkey` parameter)

Thus, to accept a private exchange proposition referred by the index `1` from the passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f`, 
the passport owner should execute the following command using it's Ethereum private key stored in file `pass_owner.key`:

```bash
./bin/privatedata-exchange accept \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 1 \
  --ownerkey ./pass_owner.key \
  --backendurl https://ropsten.infura.io
```

This is how the output looks like:

```
WARN [05-16|14:01:18.257] Reading ephemeral public key from IPFS   hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=public_key
WARN [05-16|14:01:18.773] Accepting private data exchange          exchange_index=1 encrypted_key="[12 201 213 144 184 166 215 0 173 4 39 228 62 198 102 75 221 245 254 31 183 117 86 152 34 68 218 152 49 76 233 214]"
WARN [05-16|14:01:19.771] Waiting for transaction                  hash=0x51fb2ae16e6796bcdb6025b468ffcb3d045fc7ced359db88f64ee1cb7e0ab4dd
WARN [05-16|14:01:32.305] Transaction successfully mined           tx_hash=0x51fb2ae16e6796bcdb6025b468ffcb3d045fc7ced359db88f64ee1cb7e0ab4dd gas_used=81836
```

#### Reading private data after private data exchange acceptance

After a private data exchange proposition is accepted, the data requester can read the private data by providing the following parameters:

* the passport address (`--passportaddr` parameter)
* the private data exchange index (`--exchidx` parameter)
* the name of the file with the exchange key (`--exchangekey` parameter), that was created as result of `propose` command
* the name of the file where the decrypted private data will be saved (`--datafile` parameter)

Here is the command to read private data from a passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` for the private 
data exchange referred by the index `1`, using exchange key from file `exchange.key` and writing result to standard output:

```bash
./bin/privatedata-exchange read \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 1 \
  --exchangekey ./exchange.key \
  --datafile /dev/stdout \
  --backendurl https://ropsten.infura.io
```

Below you can see how the data was read, decrypted and output to the console:

```
WARN [05-16|14:02:20.619] Reading encrypted message from IPFS      hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=encrypted_message
WARN [05-16|14:02:21.103] Reading message HMAC from IPFS           hash=QmPXKoz1jy16oHWApn5MmWgf2BcNtZTsTEAnbTtd8tw1xm filename=hmac
WARN [05-16|14:02:21.633] Writing private data to file             file_name=/dev/stdout
this is a very secret message
```

#### Closing private data exchange proposition when timed out

If the passport owner ignored the request for the private data exchange, then after 24 hours, the data requester may 
close the request and return the staked funds by calling `timeout` command.

Here's how to close the private data exchange referred by the index `2` in the passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` 
using the Ethereum private key of data requester stored in the `data_requester.key` file:

```bash
./bin/privatedata-exchange timeout \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 2 \
  --requesterkey ./data_requester.key \
  --backendurl https://ropsten.infura.io
```

The output:

```
WARN [05-17|14:01:11.231] Timeout private data exchange            exchange_index=2
WARN [05-17|14:01:11.988] Waiting for transaction                  hash=0xfa1d7f4fb5cc82c3270e123cfd2a56e0577b91d1dc56e667e66e193e5dfb57d4
WARN [05-17|14:01:41.140] Transaction successfully mined           tx_hash=0xfa1d7f4fb5cc82c3270e123cfd2a56e0577b91d1dc56e667e66e193e5dfb57d4 gas_used=23030
```

#### Closing private data exchange after acceptance

After the data requester successfully read the private data, it can confirm this by invoking the `finish` command.
When executing the command, the funds staked by the data requester and passport owner will be transferred to the passport owner.
If the data requester doesn't send the finalization request withing a predefined timespan (24 hours), the passport owner 
is allowed to finalize private data exchange, preventing the escrow being locked-up indefinitely.

This is how private data exchange referred by the index `1` in passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` may 
be finalized by the data requester using she's Ethereum private key from file `data_requester.key`:

```bash
./bin/privatedata-exchange finish \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 1 \
  --requesterkey ./data_requester.key \
  --backendurl https://ropsten.infura.io
```

The output:

```
WARN [05-16|14:03:24.275] Waiting for transaction                  hash=0x544c5b51c167efd7a1d50d26bffd706dbe2a13daefdf087afcfac6940f19b725
WARN [05-16|14:03:45.061] Transaction successfully mined           tx_hash=0x544c5b51c167efd7a1d50d26bffd706dbe2a13daefdf087afcfac6940f19b725 gas_used=31348
```

#### Opening dispute after private data exchange acceptance

If it is not possible to decrypt the data, the data requester calls the `dispute` command within 24 hours after acceptance, 
revealing the exchange key. The logic of the passport is the arbitrator who determines who the cheater is.
This is possible due to the fact that in the passport the hashes of both the data encryption key and the exchange key are saved, and
the data encryption key is XORed with the exchange key during the private data exchange acceptance by the passport owner.

When resolving a dispute, all staked funds are transferred to the side that behaved honestly.

Below you can see how the data requester tries to pretend that he could not read the private data from exchange referred 
by the index `3` from the passport `0x4026a67a2C4746b94F168bd4d082708f78d7b29f` providing valid exchange key in file `exchange3.key` and
valid Ethereum private key of data requester stored in file `data_requester.key`:

```bash
./bin/privatedata-exchange dispute \
  --passportaddr 0x4026a67a2C4746b94F168bd4d082708f78d7b29f \
  --exchidx 3 \
  --exchangekey ./exchange3.key \
  --requesterkey ./data_requester.key \
  --backendurl https://ropsten.infura.io
```

However, as a result, we see that the contract makes the only right decision that the dispute is opened unfairly 
(Dispute result: `successful=false`) and the data requester is cheater (`cheater_address=0xd2Bb3Aa3F2c0bdA6D8020f3228EabD4A89d8B951`):

```
WARN [05-16|14:38:05.921] Dispute private data exchange            exchange_index=3
WARN [05-16|14:38:06.617] Waiting for transaction                  hash=0x2770518215a20bd2d339499c2c225b430b3d6b6e10c80fb681b3e3758da92001
WARN [05-16|14:38:14.886] Transaction successfully mined           tx_hash=0x2770518215a20bd2d339499c2c225b430b3d6b6e10c80fb681b3e3758da92001 gas_used=38461
WARN [05-16|14:38:14.886] Dispute result                           successful=false cheater_address=0xd2Bb3Aa3F2c0bdA6D8020f3228EabD4A89d8B951
```

## Permissioned blockchains support

### Quorum

[Quorum](https://www.jpmorgan.com/global/Quorum)™ is an enterprise-focused version of [Ethereum](https://ethereum.org/). 
It's ideal for any application requiring high speed and high throughput processing of private transactions within a 
permissioned group of known participants.

In order to play with our SDK on Quorum network, you need to run Quorum network somewhere. The easiest way to run Quorum 
network of 7 nodes locally is by running a preconfigured Vagrant environment. Follow the 
instructions below to do this:
 
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

1. Install [Vagrant](https://www.vagrantup.com/downloads.html)

1. Download and start the Vagrant instance (note: running `vagrant up` takes approx 5 mins):
   ```
   $ git clone https://github.com/jpmorganchase/quorum-examples
   $ cd quorum-examples
   $ vagrant up
   $ vagrant ssh
   ```
   After executing these commands, you will be inside a virtual machine with all the tools to start the Quorum network.
   
   ***NOTE***: To shutdown the Vagrant instance later, run `vagrant suspend`. To delete it, run `vagrant destroy`. 
   To start from scratch, run `vagrant up` after destroying the instance. (you should run all `vagrant` commands from 
   the host machine, not from the virtual machine)
   
1. Once inside the virtual machine, run the blockchain nodes using Raft consensus:
   ```
   $ cd quorum-examples/7nodes/
   $ ./raft-init.sh
   $ ./raft-start.sh
   ```
   Make sure 7 processes of `geth` are up and running by executing `ps aux | grep geth` command.
   
   Genesis block contains 5 addresses, each of which has 1000000000 ETH:
   
   | Address                                    | Private key                                                      |
   |--------------------------------------------|------------------------------------------------------------------|
   | 0xed9d02e382b34818e88B88a309c7fe71E65f419d | e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1 |
   | 0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e | 4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9 |
   | 0x0fBDc686b912d7722dc86510934589E0AAf3b55A | 61dced5af778942996880120b303fc11ee28cc8e5036d2fdff619b5675ded3f0 |
   | 0x9186eb3d20Cbd1F5f992a950d808C4495153ABd5 | 794392ba288a24092030badaadfee71e3fa55ccef1d70c708baf55c07ed538a8 |
   | 0x0638E1574728b6D862dd5d3A3E0942c3be47D996 | 30bee17b2b8b1e774115f785e92474027d45d900a12a9d5d99af637c2d1a61bd |
   
1. When all nodes are up and running it's safe to exit from virtual machine and start reputation layer bootstrap. Run `exit`, to leave Vagrant environment:

   ```
   $ exit
   ```
   
   Now you're on the host machine.
   Vagrant environment exposes ports 22000-22007, on which Ethereum JSON RPC is available.
   You can check it's working by running command:
   
   ```
   $ curl -H "Content-Type: application/json" \
     -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
     http://localhost:22000
   ```
   
   You should see the output:
   
   ```
   {"jsonrpc":"2.0","id":1,"result":"0x0"}
   ```

1. Now follow [Building the source](#building-the-source) steps to build the full suite of reputation SDK utilities, 
   if you haven't done it yet. Use private keys from the table above and specify one of the Quorum node (like `http://localhost:22000`) 
   as `-backendurl` parameter to make transactions.
   
   For example:
   ```
   $ echo e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1 > bootstrap_owner.key
   $ echo 4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9 > passport_owner.key
   $ echo 61dced5af778942996880120b303fc11ee28cc8e5036d2fdff619b5675ded3f0 > data_provider.key
   $ ./bin/deploy-bootstrap -ownerkey ./bootstrap_owner.key -backendurl http://localhost:22000
   $ ./bin/deploy-passport -ownerkey ./passport_owner.key \
        -factoryaddr 0x9d13C6D3aFE1721BEef56B55D303B09E021E27ab \
        -backendurl http://localhost:22000
   $ echo Johny | ./bin/write-fact -fkey name -ftype string \
        -ownerkey ./data_provider.key \
        -passportaddr 0x4AEb3678b689DbB3F502D927580f9829001C4BB6 \
        -backendurl http://localhost:22000
   $ ./bin/read-fact -out /dev/stdout -fkey name -ftype string \
        -factprovideraddr 0x0fBDc686b912d7722dc86510934589E0AAf3b55A \
        -passportaddr 0x4AEb3678b689DbB3F502D927580f9829001C4BB6 \
        -backendurl http://localhost:22000
   ...
   ```