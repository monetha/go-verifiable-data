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

### Contracts update

After Ethereum contracts code is updated and artifacts are created:
1. Copy all artifacts to [contracts/code](contracts/code) folder.
1. Run `go generate` command in [contracts](contracts) folder to convert Ethereum contracts into Go package.
1. Commit new/updated `.go` files.