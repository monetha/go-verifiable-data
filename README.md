# protocol-go-sdk

## Development

### Build

Install dependencies:

```bash
make dependencies
```

Make sure that the tests and the linters pass:

```bash
make lint
make test
```

To get the test coverage report for all packages, run the command:

```bash
make cover
```

Тest coverage results (`cover.out`, `cover.html`) will be put in `./.cover` directory.

### Contracts update

After Ethereum contracts code is updated and artifacts are created:
1. Copy all artifacts to [`contracts/code`](contracts/code) folder.
1. Run `go generate` command in [`contracts`](contracts) folder to convert Ethereum contracts into Go package.
1. Commit new/updated files.

### Executables

The protocol-go-sdk project comes with several executables found in the [`cmd`](cmd) directory.

| Command    | Description |
|:----------:|-------------|
| [`deploy-bootstrap`](cmd/deploy-bootstrap) | Utility tool to deploy three contracts at once: [PassportLogic](contracts/code/PassportLogic.sol), [PassportLogicRegistry](contracts/code/PassportLogicRegistry.sol), [PassportFactory](contracts/code/PassportFactory.sol). |
| [`deploy-passport`](cmd/deploy-passport) | Utility tool to deploy [Passport](contracts/code/Passport.sol) contracts using already deployed [PassportFactory](contracts/code/PassportFactory.sol). |
| [`write-fact`](cmd/write-fact) | Utility tool to write facts to passport. |
| [`read-fact`](cmd/read-fact) | Utility tool to read facts from passport. |
| [`passport-list`](cmd/passport-list) | Utility tool for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract. |
