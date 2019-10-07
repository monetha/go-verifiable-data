# Fact provider registry tool

## Usage

`./provider-registry [OPTIONS] <COMMAND> [COMMAND-OPTIONS]`

See commands below for individual usage or access command line help by running:

`./provider-registry --help`

and for specific command:

`./provider-registry <COMMAND> --help`

To output version of tool, run:

`./provider-registry --version`

# Commands

## deploy

Command to deploy [FactProviderRegistry](../../contracts/code/FactProviderRegistry.sol) contract.

### Usage

```
Usage:
  provider-registry [OPTIONS] deploy [deploy-OPTIONS]

Application Options:
  -v, --version                Print the version of tool and exit

Help Options:
  -h, --help                   Show this help message

[deploy command options]
          --quorum_enclave=    Quorum enclave url for private transactions
          --quorum_privatefor= Quorum nodes public keys to make transaction private for, separated by commas
          --ownerkey=          fact provider registry owner private key filename
          --backendurl=        Ethereum backend URL
          --verbosity=         log verbosity (0-9) (default: 2)
          --vmodule=           log verbosity pattern
```

### Examples

* Deploying [FactProviderRegistry](../../contracts/code/FactProviderRegistry.sol) contract in Ropsten network using Ethereum private key stored in file `./owner.key`:
  ```bash
  ./provider-registry deploy --ownerkey ./owner.key --backendurl https://ropsten.infura.io
  ```

## set

Command to set/add fact provider information

### Usage

```
Usage:
  provider-registry [OPTIONS] set [set-OPTIONS]

Application Options:
  -v, --version                Print the version of tool and exit

Help Options:
  -h, --help                   Show this help message

[set command options]
          --quorum_enclave=    Quorum enclave url for private transactions
          --quorum_privatefor= Quorum nodes public keys to make transaction private for, separated by commas
          --ownerkey=          fact provider registry owner private key filename
          --backendurl=        Ethereum backend URL
          --registryaddr=      Ethereum address of fact provider registry contract
          --provideraddr=      Ethereum address of fact provider
          --providername=      Name of fact provider
          --providerpassaddr=  Ethereum address of passport of fact provider
          --providerwebsite=   Website of fact provider
          --verbosity=         log verbosity (0-9) (default: 2)
          --vmodule=           log verbosity pattern
```

### Examples

* Add fact provider information (fact provider address `0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278`, name: `Monetha`, 
  passport address: `0x556e0345ebaa820409d4ac35aded6e54fb8bbf27`, website: `https://www.monetha.io`) to the registry 
  (address `0xf9dbc37bbdc68e0ba03185f1877059c595dcf083`) in Ropsten network using Ethereum private key stored 
  in file `./owner.key`:
  ```
  ./provider-registry set \
      --ownerkey ./owner.key \
      --backendurl https://ropsten.infura.io \
      --registryaddr 0xf9dbc37bbdc68e0ba03185f1877059c595dcf083 \
      --provideraddr 0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278 \
      --providername Monetha \
      --providerpassaddr 0x556e0345ebaa820409d4ac35aded6e54fb8bbf27 \
      --providerwebsite https://www.monetha.io
  ```

## get

Command to get fact provider information.

### Usage

```
Usage:
  provider-registry [OPTIONS] get [get-OPTIONS]

Application Options:
  -v, --version                Print the version of tool and exit

Help Options:
  -h, --help                   Show this help message

[get command options]
          --quorum_enclave=    Quorum enclave url for private transactions
          --quorum_privatefor= Quorum nodes public keys to make transaction private for, separated by commas
          --backendurl=        Ethereum backend URL
          --registryaddr=      Ethereum address of fact provider registry contract
          --provideraddr=      Ethereum address of fact provider
          --verbosity=         log verbosity (0-9) (default: 2)
          --vmodule=           log verbosity pattern
```

### Examples

* Get fact provider information (fact provider address `0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278`) from the the registry 
  (address `0xf9dbc37bbdc68e0ba03185f1877059c595dcf083`) in Ropsten network:
  ```
  ./provider-registry get \
      --backendurl https://ropsten.infura.io \
      --registryaddr 0xf9dbc37bbdc68e0ba03185f1877059c595dcf083 \
      --provideraddr 0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278
  ```

## delete

Command to delete fact provider information

### Usage

```
Usage:
  provider-registry [OPTIONS] delete [delete-OPTIONS]

Application Options:
  -v, --version                Print the version of tool and exit

Help Options:
  -h, --help                   Show this help message

[delete command options]
          --quorum_enclave=    Quorum enclave url for private transactions
          --quorum_privatefor= Quorum nodes public keys to make transaction private for, separated by commas
          --ownerkey=          fact provider registry owner private key filename
          --backendurl=        Ethereum backend URL
          --registryaddr=      Ethereum address of fact provider registry contract
          --provideraddr=      Ethereum address of fact provider
          --verbosity=         log verbosity (0-9) (default: 2)
          --vmodule=           log verbosity pattern
```

## Examples

* Delete fact provider information (fact provider address `0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278`) from the the registry 
  (address `0xf9dbc37bbdc68e0ba03185f1877059c595dcf083`) in Ropsten network using Ethereum private key stored in file `./owner.key`:
  ```
  ./provider-registry delete \
      --ownerkey ./owner.key \
      --backendurl https://ropsten.infura.io \  
      --registryaddr 0xf9dbc37bbdc68e0ba03185f1877059c595dcf083 \
      --provideraddr 0x7d05b13c74b173fa77b0de73baec1e8f8d2b9278
  ```