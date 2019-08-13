# passport-permission

A utility tool that allows a passport holder to allow or deny a fact provider to write/delete facts to/from a passport.
By default any fact provider can write to a passport, but a passport holder can change permissions that allow only 
fact providers from the whitelist to write to a passport.

## Usage

Usage of `./passport-permission`:
```
  -add value
    	add fact provider to the whitelist
  -backendurl string
    	backend URL (simulated backend used if empty)
  -onlywhitelist value
    	enables or disables the use of a whitelist of fact providers
  -ownerkey string
    	owner private key filename
  -ownerkeyhex string
    	private key as hex (for testing)
  -passaddr string
    	Ethereum address of passport contract
  -remove value
    	remove fact provider from the whitelist
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
  -quorum_privatefor
        Quorum nodes public keys to make transaction private for, separated by commas
  -quorum_enclave
        Quorum enclave url for private transactions
```

## Examples

* Owner of a passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
adds fact provider `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

    ```bash
    ./passport-permission -ownerkey pass_owner.key \
      -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -add 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      -backendurl https://ropsten.infura.io
    ```
    
* Owner of a passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
removes fact provider `0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d` to the whitelist in Ropsten network:

    ```bash
    ./passport-permission -ownerkey pass_owner.key \
      -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -remove 0x5b2AE3b3A801469886Bb8f5349fc3744cAa6348d \
      -backendurl https://ropsten.infura.io
    ```
    
* Owner of a passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows to store the facts only to fact providers from the whitelist:

    ```bash
    ./passport-permission -ownerkey pass_owner.key \
      -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -onlywhitelist true \
      -backendurl https://ropsten.infura.io
    ```
    
* Owner of a passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
allows any fact provider to write the facts:

    ```bash
    ./passport-permission -ownerkey pass_owner.key \
      -passaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -onlywhitelist false \
      -backendurl https://ropsten.infura.io
    ```