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

* Read the entire change history for the passport [`0x9CfabB3172DFd5ED740c3b8A327BF573226c5064`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
  in `Ropsten` block-chain and write it to the file `/dev/stdout` (outputs to the screen):
    ```bash
    ./read-history -out /dev/stdout \
      -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
      -backendurl https://ropsten.infura.io
    ```
* Retrieve the history value of type `txdata` stored in transaction [`0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a`](https://ropsten.etherscan.io/tx/0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a) from passport 
    [`0x9CfabB3172DFd5ED740c3b8A327BF573226c5064`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./history_image.jpg`:
    ```bash
    ./read-history -out history_image.jpg \
      -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
      -ftype txdata \
      -txhash 0x627913f620990ec12360a6f1fda4887ea837b41e2f6cbae90e24322dc8cf8b1a \
      -backendurl https://ropsten.infura.io
    ```

* Retrieve the history value of type `ipfs` stored in transaction [`0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f`](https://ropsten.etherscan.io/tx/0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f) from passport 
    [`0x9CfabB3172DFd5ED740c3b8A327BF573226c5064`](https://ropsten.etherscan.io/address/0x9cfabb3172dfd5ed740c3b8a327bf573226c5064)
    in `Ropsten` block-chain and write it to the file `./Monetha_WP.pdf`:
    ```bash
    ./read-history -out Monetha_WP.pdf \
      -passportaddr 0x9CfabB3172DFd5ED740c3b8A327BF573226c5064 \
      -ftype ipfs \
      -txhash 0x91c5d11c7f220660fb2c98273627e9c2f01b59e32163c760a4a9a836f7758f7f \
      -backendurl https://ropsten.infura.io
    ```