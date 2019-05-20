# passport-list

Utility tool for getting a list of passports created using specific [PassportFactory](../../contracts/code/PassportFactory.sol) contract.

## Usage

Usage of `./passport-list`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -factoryaddr string
    	Ethereum address of passport factory contract
  -out string
    	save retrieved passports to the specified file
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
```

## Examples

* Get example passport from the simulated backend:
  ```bash
  ./passport-list
  ```

* Get all passports created by the passport factory [`0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2`](https://ropsten.etherscan.io/address/0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2#code) in Ropsten network
  and write them to the file `./passports.csv`:
  ```bash
  ./passport-list -out ./passports.csv \
    -backendurl https://ropsten.infura.io \
    -factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
  ```
