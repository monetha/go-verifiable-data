# protocol-go-sdk

* [Building the source](#building-the-source)
    * [Prerequisites](#prerequisites)
    * [Build](#build)
    * [Executables](#executables)
* [Contributing](#contributing)
    * [Making changes](#making-changes)
    * [Contracts update](#contracts-update)
    * [Formatting source code](#formatting-source-code)
* [Bootstrap reputation protocol](#bootstrap-reputation-protocol)
* [Usage](#usage)
    * [Deploying passport](#deploying-passport)
    * [Passport list](#passport-list)
    * [Writing facts](#writing-facts)
    * [Reading facts](#reading-facts)
    * [Changing passport permissions](#changing-passport-permissions)
    * [Reading facts history](#reading-facts-history)

## Building the source

### Prerequisites

1. Make sure you have [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed.
1. Install [Go 1.11](https://golang.org/dl/)
1. Setup `$GOPATH` environment variable as described [here](https://github.com/golang/go/wiki/SettingGOPATH).
1. Clone the repository:
    ```bash
    mkdir -p $GOPATH/src/gitlab.com/monetha
    cd $GOPATH/src/gitlab.com/monetha
    git clone git@gitlab.com:monetha/protocol-go-sdk.git
    cd protocol-go-sdk
    ```

### Build

Install dependencies:

    make dependencies

Once the dependencies are installed, run 

    make cmd

to build the full suite of utilities. After the executable files are built, they can be found in the directory `./bin/`.

### Executables

The protocol-go-sdk project comes with several executables found in the [`cmd`](cmd) directory.

| Command    | Description |
|:----------:|-------------|
| [`deploy-bootstrap`](cmd/deploy-bootstrap) | Utility tool to deploy three contracts at once: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol). |
| [`deploy-passport`](cmd/deploy-passport) | Utility tool to deploy [Passport](contracts/code/Passport.sol) contracts using already deployed [PassportFactory](contracts/code/PassportFactory.sol). |
| [`write-fact`](cmd/write-fact) | Utility tool to write facts to passport. |
| [`read-fact`](cmd/read-fact) | Utility tool to read facts from passport. |
| [`passport-list`](cmd/passport-list) | Utility tool for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract. |
| [`passport-permission`](cmd/passport-permission) | Utility tool that allows a passport holder to allow or deny a fact provider to write/delete facts to/from a passport. By default any fact provider can write to a passport, but a passport holder can change permissions that allow only fact providers from the whitelist to write to a passport. |
| [`read-history`](cmd/read-history) | Utility tool for reading the history of passport changes. |

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

## Bootstrap reputation protocol

Monetha has already deployed this set of auxiliary reputation protocol contracts on Ropsten test network. The contract addresses deployed on Ropsten:

| Contract      | Address                                      |
|---------------|----------------------------------------------|
| `PassportLogic` | [`0x4FBF5019E0B7B2470810e67E10aAA75A57319a9b`](https://ropsten.etherscan.io/address/0x4FBF5019E0B7B2470810e67E10aAA75A57319a9b) |
| `PassportLogicRegistry`  | [`0xabA015Fc83E9B88e8334bD9b536257db75e05295`](https://ropsten.etherscan.io/address/0xabA015Fc83E9B88e8334bD9b536257db75e05295) |
| `PassportFactory` | [`0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`](https://ropsten.etherscan.io/address/0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2) |

Consider the process of deploying your own set of auxiliary repoutation protocol contracts to experiment with our implementation. If you are going to deploy your contracts, then you will have to support them yourself.

This means that if the reputation protocol logic of the passport is updated by Monetha developers, you'll need to deploy a new `PassportLogic` contract, register it 
in an existing `PassportLogicRegistry` contract (by calling `addPassportLogic` method) and finally make it active (by calling `setCurrentPassportLogic`).

If you use a set of Monetha deployed reputation protocol contracts, then the reputation passport logic is always up-to-date with latest fixes and features.

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

In order to create a passport and start using it, you need to use auxiliary reputation protocol contracts: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol).

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

After the data has been read from the Ethereum blockchain and written to the file `./fact_image.jpg`, try to open the image..

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
WARN [10-31|11:03:53.291] Loaded configuration                     backend_url=https://ropsten.infura.io passport=0x9CfabB3172DFd5ED740c3b8A327BF573226c5064
fact_provider,key,data_type,change_type,block_number,tx_hash
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,TxData,Updated,4177015,0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a
0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d,monetha.jpg,TxData,Updated,4337297,0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8
WARN [10-31|11:03:54.643] Done.
```

The CSV output can be saved to a file and converted to the table:

| fact_provider | key | data_type | change_type | block_number | tx_hash |
|---------------|-----|-----------|-------------|--------------|---------|
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | TxData | Updated | 4177015 | 0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a |
| 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d | monetha.jpg | TxData | Updated | 4337297 | 0x31e06af4e04450333d468835c995fc02622c1b07ae0feeb4c7afe73c5a2e3ed8 |

As we can see, there were only two fact updates of type `TxData` (under the same key `monetha.jpg`) by the same data provider `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d`.
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