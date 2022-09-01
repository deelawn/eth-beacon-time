# eth-beacon-time

Provide a slot to get the time of that slot or a unix time to get the slot at that time.

`chain-id` is an optional parameter and the current valid values are `mainnet` and `prater`; `mainnet` is the default.

Usage:
```
❯ go run ./cmd/eth-beacon-time --slot 4000000

2022-06-10 03:20:00 +0200 CEST

❯ go run ./cmd/eth-beacon-time --time 1652565600

3811800
```