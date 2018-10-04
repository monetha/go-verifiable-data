# write-fact

Utility tool to write facts to passport.

## tx data vs. string

Cumulative gas used in simulated backend to store bytes:

| Number of bytes  |     tx data    |
|-----------------:|---------------:|
| 10 | 70436 |
| 100 | 76598 |
| 500 | 103870 |
| 1000 | 138016 |
| 5000 | 410814 |
| 10000 | 751864 |
| 50000 | 3483963 |
| 100000 | 6907662 |
| 110000 | 7593621 |
| 120000 | 8279814 |
| 130000 | 8966537 |

## Examples

* Store 1000 characters of `a` under the key `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa` in simulated backend 
using fact provider private key `1dae9ab9e0c080371c56d816f4b6323e6c229e1cea4d15bc7f828c40ad9729d6`:
  ```bash
  head -c 1000 < /dev/zero | tr '\0' '\141' | ./write-fact \
    -factkey aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa \
    -ownerkeyhex 1dae9ab9e0c080371c56d816f4b6323e6c229e1cea4d15bc7f828c40ad9729d6
  ```
