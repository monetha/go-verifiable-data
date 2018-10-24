# protocol-go-sdk

* [Building the source](#building-the-source)
    * [Prerequisites](#prerequisites)
    * [Build](#build)
    * [Executables](#executables)
* [Contributing](#contributing)
    * [Making changes](#making-changes)
    * [Contracts update](#contracts-update)
    * [Formatting source code](#formatting-source-code)
* [Usage](#usage)

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

## Usage