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
| `PassportLogic` | [`0xf7adefec07440c9846afe5cc7ecca6821a831208`](https://ropsten.etherscan.io/address/0xf7adefec07440c9846afe5cc7ecca6821a831208) |
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