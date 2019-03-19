# read-history

Utility tool for reading the history of passport changes.

## Usage

Usage of `./read-history`:
```
  -backendurl string
    	backend URL (simulated backend used if empty)
  -ftype string
    	the data type of fact (txdata, string, bytes, address, uint, int, bool, ipfs)
  -ipfsurl string
    	IPFS node address (default "https://ipfs.infura.io:5001")
  -out string
    	save retrieved data to the specified file
  -passportaddr value
    	Ethereum address of passport contract
  -txhash value
    	the transaction hash to read history value from
  -verbosity int
    	log verbosity (0-9) (default 2)
  -vmodule string
    	log verbosity pattern
```

When the `-ftype` and` -txhash` options are omitted, the entire change history is returned. 

Specifying the parameters `-ftype` and` -txhash` allows you to read the value of the specified type from the specified transaction.

## Examples

* Read the entire change history for the passport [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen):
    ```bash
    ./read-history -out /dev/stdout \
      -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -backendurl https://ropsten.infura.io
    ```
* Retrieve the history value of type `txdata` stored in transaction [`0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632`](https://ropsten.etherscan.io/tx/0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632) from passport 
    [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./history_image.jpg`:
    ```bash
    ./read-history -out history_image.jpg \
      -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -ftype txdata \
      -txhash 0xd43201d6b23a18b90a53bf7ef1fffad0b04af603c039b6617601a225a129c632 \
      -backendurl https://ropsten.infura.io
    ```

* Retrieve the history value of type `ipfs` stored in transaction [`0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca`](https://ropsten.etherscan.io/tx/0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca) from passport 
    [`0x1C3A76a9A27470657BcBE7BfB47820457E4DB682`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./Monetha_WP.pdf`:
    ```bash
    ./read-history -out Monetha_WP.pdf \
      -passportaddr 0x1C3A76a9A27470657BcBE7BfB47820457E4DB682 \
      -ftype ipfs \
      -txhash 0xbc8a86f54a467edbec32fbf27c08e7077221dd69bbea79707889ac6f787fe0ca \
      -backendurl https://ropsten.infura.io
    ```