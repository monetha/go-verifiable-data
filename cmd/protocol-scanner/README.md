# protocol-scanner

This application serves static content of a WebAssembly module which works completely in a browser and uses publicly available Ethereum block-chain 
nodes to provide the following data:

* the list of deployed passports
* the history of passport changes

WebAssembly is produced from the [Go source code](web/main.go) that actively uses Go SDK of reputation protocol.

This is made possible by the recently added [an experimental port to WebAssembly in Go 1.11](https://github.com/golang/go/wiki/WebAssembly),
and also by [our changes](https://github.com/ethereum/go-ethereum/pull/17709) that were accepted by the [go-ethereum](https://github.com/ethereum/go-ethereum) community.

## Usage

Usage of `./protocol-scanner`:
```
  -listen string
    	listen address (default ":8080")
```

By default, the application serves content on port `8080`, but the port can be changed by specifying the `-listen` parameter.
For example, you can change the port to `8081` using the command:

    ./protocol-scanner -listen :8081

If you run the program with the default settings, then after launch you just need to open the following link in the browser:
[http://localhost:8080](http://localhost:8080)

In your browser make sure you use the correct parameters for `Backend URL` and `Passport factory address`.
By default, these parameters are set to the following values:

* `Backend URL`: `https://ropsten.infura.io`
* `Passport factory address`: `0x87b7Ec2602Da6C9e4D563d788e1e29C064A364a2`