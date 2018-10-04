# write-fact

Utility tool to write facts to passport.

## tx data vs. string

Cumulative gas usage in simulated backend to store number of character of `a` under the key 
`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa` using different data types:

| Number of characters |     `txdata`, gas used    |  `bytes`, gas used |  `string`, gas used |
|---------------------:|--------------------------:|-------------------:|-------------------:|
| 10 | 70436 | 71079 | 71211 |
| 100 | 76598 | 157571 | 157703 |
| 500 | 103870 | 425756 | 425888 |
| 1000 | 138016 | 781119 | 781251 |
| 5000 | 410814 | 3563467 | 3563599 |
| 10000 | 751864 | 7036521 | 7036653 |
| 50000 | 3483963 | - | - |
| 100000 | 6907662 | - | - |
| 110000 | 7593621 | - | - |
| 120000 | 8279814 | - | - |
| 130000 | 8966537 | - | - |

## Examples

* Store 1000 characters of `a` under the key `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa` as `txdata` in simulated backend 
using fact provider private key `1dae9ab9e0c080371c56d816f4b6323e6c229e1cea4d15bc7f828c40ad9729d6`:
  ```bash
  head -c 1000 < /dev/zero | tr '\0' '\141' | ./write-fact \
    -fkey aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa \
    -ftype txdata \
    -ownerkeyhex 1dae9ab9e0c080371c56d816f4b6323e6c229e1cea4d15bc7f828c40ad9729d6
  ```
