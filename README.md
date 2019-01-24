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
1. Install [Go 1.11](https://golang.org/dl/)
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
  golang:1.11 \
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
| `PassportLogic` | [`0x0361a024040E7020251fF0756Bb40B8e136B9c9f`](https://ropsten.etherscan.io/address/0x0361a024040E7020251fF0756Bb40B8e136B9c9f) |
| `PassportLogicRegistry`  | [`0xabA015Fc83E9B88e8334bD9b536257db75e05295`](https://ropsten.etherscan.io/address/0xabA015Fc83E9B88e8334bD9b536257db75e05295) |
| `PassportFactory` | [`0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2) |

The contract addresses deployed on Mainnet:

| Contract      | Address                                      |
|---------------|----------------------------------------------|
| `PassportLogic` | [`0xbCd4C9ba1EfB413b1AC952EfaA2374F98641eb7f`](https://etherscan.io/address/0xbCd4C9ba1EfB413b1AC952EfaA2374F98641eb7f) |
| `PassportLogicRegistry`  | [`0x3dC70507087D36A726a0A3fD99eb2D4b513248B0`](https://etherscan.io/address/0x3dC70507087D36A726a0A3fD99eb2D4b513248B0) |
| `PassportFactory` | [`0x9F58301392696aaa0A23FBA7B8dE3118c72A8685`](https://etherscan.io/address/0x9F58301392696aaa0A23FBA7B8dE3118c72A8685) |

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
WARN [09-19|14:52:47.640] Loaded configuration                     owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 backend_url=https://ropsten.infura.io
WARN [09-19|14:52:47.640] Checking balance                         address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [09-19|14:52:48.675] Deploying PassportLogic                  owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55
WARN [09-19|14:52:49.426] Waiting for transaction                  hash=0x5a6e6c4110852149113619905ac8752b9120072082f525c8ad84b81c99d53ccb
WARN [09-19|14:53:06.042] PassportLogic deployed                   contract_address=0x977De088AD659D37c064DB4dc2738ACf3aE09dd8
WARN [09-19|14:53:06.042] Deploying PassportLogicRegistry          owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 impl_version=0.1 impl_address=0x977De088AD659D37c064DB4dc2738ACf3aE09dd8
WARN [09-19|14:53:06.819] Waiting for transaction                  hash=0x984c87fb83607c2b1e8bb008e5d6ab07429149c817f2ce368901136b0a786902
WARN [09-19|14:53:10.965] PassportLogicRegistry deployed           contract_address=0xabA015Fc83E9B88e8334bD9b536257db75e05295
WARN [09-19|14:53:10.965] Deploying PassportFactory                owner_address=0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55 registry=0xabA015Fc83E9B88e8334bD9b536257db75e05295
WARN [09-19|14:53:11.585] Waiting for transaction                  hash=0x2467e6fa520fbc09aab94758a766a891aa3222d56ab56c28783a3c93e29a892e
WARN [09-19|14:53:15.726] PassportFactory deployed                 contract_address=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
WARN [09-19|14:53:15.726] Done.
```

In the output you can find the addresses of all the deployed contracts.


## Usage

In order to create a passport and start using it, you need to use auxiliary reputation layer contracts: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol).

### Deploying passport

Before creating a passport for a specific Ethereum address, store the private key of this Ethereum address in the file `pass_owner.key`.
Make sure that the passport owner has enough money to create a passport contract. Usually passport contract deployment takes `425478` gas.

To create a passport contract you need to know address of the `PassportFactory` contract. Let's try to create a passport in Ropsten
using the `PassportFactory` contract deployed by Monetha ([`0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2)):

```bash
./bin/deploy-passport -ownerkey ./pass_owner.key \
  -factoryaddr 0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2 \
  -backendurl https://ropsten.infura.io
```

Below you can see the output of the command to create a passport for the address `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d`.

```
WARN [10-24|16:19:32.928] Loaded configuration                     owner_address=0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d backend_url=https://ropsten.infura.io factory=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
WARN [10-24|16:19:33.971] Getting balance                          address=0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d
WARN [10-24|16:19:34.138] Initializing PassportFactory contract    factory=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
WARN [10-24|16:19:34.139] Deploying Passport contract              owner_address=0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d
WARN [10-24|16:19:34.823] Waiting for transaction                  hash=0x96f4c583994c2d1c033a0722f7cbe8d85c636b62d1f5fcd8bb0b32346c61c4a9
WARN [10-24|16:19:47.364] Transaction successfully mined           tx_hash=0x96f4c583994c2d1c033a0722f7cbe8d85c636b62d1f5fcd8bb0b32346c61c4a9 cumulative_gas_used=1445116
WARN [10-24|16:19:47.365] Passport deployed                        contract_address=0x86eEb0D360D286BcF9211780878fe0D0c0e3fF00
WARN [10-24|16:19:47.365] Initializing Passport contract           passport=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
WARN [10-24|16:19:47.365] Claiming ownership                       owner_address=0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d
WARN [10-24|16:19:48.141] Waiting for transaction                  hash=0x1aa26e7416e97747f0da7103cfc1dd3662bd3bbee3dee4680f378e13057a09b7
WARN [10-24|16:20:04.819] Transaction successfully mined           tx_hash=0x1aa26e7416e97747f0da7103cfc1dd3662bd3bbee3dee4680f378e13057a09b7 cumulative_gas_used=1930990
WARN [10-24|16:20:04.825] Done.
```

As you can see in the line `Passport deployed`, a passport contract was created at address [`0x86eEb0D360D286BcF9211780878fe0D0c0e3fF00`](https://ropsten.etherscan.io/address/0x86eEb0D360D286BcF9211780878fe0D0c0e3fF00).

### Passport list

The passport factory allows you to get a list of all the passports that have been created.

Let's try to get a list of all passports using the address of `PassportFactory` contract deployed by Monetha ([`0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2))
in Ropsten network:

```bash
./bin/passport-list -out /dev/stdout \
  -backendurl https://ropsten.infura.io \
  -factoryaddr 0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
```

You should see something like this:

```
WARN [10-24|16:30:43.549] Loaded configuration                     factory_provider=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2 backend_url=https://ropsten.infura.io
WARN [10-24|16:30:44.670] Initialising passport factory contract   passport_factory=0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2
WARN [10-24|16:30:45.000] Writing collected passports to file
passport_address,first_owner,block_number,tx_hash
0x9CfabB3172DFd5ED740c3b8A327BF573226c5064,0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55,4105235,0x5a26791f5404f7d26c9c75e4fa006d851162f4bbaacc49372ce45d89db8fd967
0x2ff877C92458F995332bc189F258eF8fB8458050,0xA12eB9Cde44664B6513D66f1fc4d43c951d4594e,4276542,0x639262c4abf2868e376e6b08baa5663a2449b18fc668836b5451d07f24c04db5
0x86eEb0D360D286BcF9211780878fe0D0c0e3fF00,0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,4292633,0x96f4c583994c2d1c033a0722f7cbe8d85c636b62d1f5fcd8bb0b32346c61c4a9
WARN [10-24|16:30:45.000] Done.
```

The output can be saved to a file and converted to the table. Currently three passports are deployed:

|passport_address|first_owner|block_number|tx_hash|
|----------------|-----------|------------|-------|
|0x9CfabB3172DFd5ED740c3b8A327BF573226c5064|0xDdD9b3Ea9d65cfD12b18ceA4E6f7Df4948ec4C55|4105235|0x5a26791f5404f7d26c9c75e4fa006d851162f4bbaacc49372ce45d89db8fd967|
|0x2ff877C92458F995332bc189F258eF8fB8458050|0xA12eB9Cde44664B6513D66f1fc4d43c951d4594e|4276542|0x639262c4abf2868e376e6b08baa5663a2449b18fc668836b5451d07f24c04db5|
|0x86eEb0D360D286BcF9211780878fe0D0c0e3fF00|0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d|4292633|0x96f4c583994c2d1c033a0722f7cbe8d85c636b62d1f5fcd8bb0b32346c61c4a9|

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
`0x9CfabB3172DFd5ED740c3b8A327BF573226c5064`:

```bash
./bin/write-fact -ownerkey fact_provider.key \
  -fkey monetha.jpg \
  -ftype txdata \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
```

Similarly, you can store the same image from the file `~/Downloads/monetha.jpg` under the key `monetha.jpg` as `ipfs` in the same passport.
Keep in mind, the data will be stored in IPFS, only IPFS hash will be stored in the Ethereum storage:

```bash
./bin/write-fact -ownerkey fact_provider.key \
  -fkey monetha.jpg \
  -ftype ipfs \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -backendurl https://ropsten.infura.io < ~/Downloads/monetha.jpg
```

### Reading facts

After the fact provider has written the public data to the passport, the data can be read by anyone.
To read the data you need to know: the address of the passport, the address of the fact provider who stored the data, 
the key under which the data was stored and the type of data.

Let's try to retrieve image from passport `0x9CfabB3172DFd5ED740c3b8A327BF573226c5064` that was stored by the fact provider
`0x5b2ae3b3a801469886bb8f5349fc3744caa6348d` under the key `monetha.jpg` as `txdata` data type and write it to the file 
`./fact_image.jpg`:

```bash
./bin/read-fact -out ./fact_image.jpg \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -factprovideraddr 0x5b2ae3b3a801469886bb8f5349fc3744caa6348d \
  -fkey monetha.jpg \
  -ftype txdata \
  -backendurl https://ropsten.infura.io
```

After the data has been read from the Ethereum blockchain and written to the file `./fact_image.jpg`, try to open the image.

To get the same file that was previously saved in IPFS, just change the parameter `-ftype` to `ipfs`:

```bash
./bin/read-fact -out ./ipfs_fact_image.jpg \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
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

Consider an example of how owner of a passport `0x9CfabB3172DFd5ED740c3b8A327BF573226c5064` adds fact provider 
`0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

```bash
./bin/passport-permission -ownerkey pass_owner.key \
  -passaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -add 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
  -backendurl https://ropsten.infura.io
```

Please note that the passport owner’s private key is stored in the file `pass_owner.key`.

After executing the command, any fact provider is still allowed to store the facts in the passport. Let's fix it!

Owner of a passport `0x9CfabB3172DFd5ED740c3b8A327BF573226c5064` may allow to store the facts only to fact providers 
from the whitelist by running the command:

```bash
./bin/passport-permission -ownerkey pass_owner.key \
  -passaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -onlywhitelist true \
  -backendurl https://ropsten.infura.io
```

More examples can be found [here](cmd/passport-permission#examples).

### Reading facts history

The SDK allows you to see the history of absolutely all changes of facts in the passport.

Let's try to retrieve the entire change history for the passport [`0x9CfabB3172DFd5ED740c3b8A327BF573226c5064`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen, but you can change this to the file name to write to the file):

```bash
./bin/read-history -out /dev/stdout \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -backendurl https://ropsten.infura.io
```

After running the command you should see something like this:

```
WARN [11-21|17:49:37.251] Loaded configuration                     backend_url=https://ropsten.infura.io passport=0x9CfabB3172DFd5ED740c3b8A327BF573226c5064
fact_provider,key,data_type,change_type,block_number,tx_hash
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,TxData,Updated,4177015,0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,TxData,Updated,4337297,0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,Monetha_WP.pdf,IPFS,Updated,4468222,0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f
WARN [11-21|17:49:38.709] Done.
```

The CSV output can be saved to a file and converted to the table:

| fact_provider | key | data_type | change_type | block_number | tx_hash |
|---------------|-----|-----------|-------------|--------------|---------|
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | TxData | Updated | 4177015 | 0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a |
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | TxData | Updated | 4337297 | 0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8 |
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | Monetha_WP.pdf | IPFS | Updated | 4468222 | 0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f |

As we can see, there were two fact updates of type `TxData` (under the same key `monetha.jpg`) by the same data provider `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d`, 
and one update of type `IPFS` by the same data provider.
The `block_number` and `tx_hash` columns allow us to understand in which block and in which transaction the changes were made.
The `change_type` column may contain either `Updated` or `Deleted` values. Even if the value of a fact has been deleted, we can read its value as it was before the deletion.

Let's read what the value of the fact was during the first update. To do this, we need to specify the transaction hash `0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a` and the type of data `txdata`:

```bash
./bin/read-history -out monetha1.jpg \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -ftype txdata \
  -txhash 0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a \
  -backendurl https://ropsten.infura.io
```

Similarly, we can read what fact value was written in the second transaction `0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8`:

```bash
./bin/read-history -out monetha2.jpg \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -ftype txdata \
  -txhash 0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8 \
  -backendurl https://ropsten.infura.io
```

Now you can compare pictures `monetha1.jpg` and `monetha2.jpg` to see what changes have been made.

To read fact value of type `TxData` in the third transaction parameter `-ftype` should be changed to `ipfs` and `-txhash` 
to `0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f`:

```bash
./bin/read-history -out Monetha_WP.pdf \
  -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
  -ftype ipfs \
  -txhash 0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f \
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

and open [http://localhost:8080](http://localhost:8080) in your browser.

More details can be found [here](cmd/passport-scanner).

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