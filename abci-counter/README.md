# ABCI Counter
This is a minimal example of a blockchain application using:
> CometBFT (consensus & networking) + a custom ABCI app (state machine)

The ABCI app here is a simple counter that:
- starts at 0
- accepts transaction with integer values
- only allows the next number in sequence(e.g., tx "1" -> ok, tx "3" -> rejected if current stat is 1)

This simulates a stateful blockchain that:
- validates transactions before block inclusion (CheckTx)
- updates application state on block commit (DeliverTx)
- Commits state hashes (Commit)

## Purposes of this project
1. Learn ABCI: application-consensus boundary
2. Minimal blockchain without cosmos sdk
3. Foundation for custom blockchains
  - template for voting system, key-value store, and a custom smart contract VM

## How to broadcast tx
Using curl
```sh
curl 'http://localhost:26657/broadcast_tx_commit?tx="1"'
```

- reference: https://github.com/cometbft/cometbft/blob/v0.37.x/rpc/core/doc.go

## How does it work
```
```
[ CometBFT Node ]
        │
        ▼
[ ABCI Socket Server ]
        │
        ▼
[ ABCI App (Go) — our counter logic ]
```
```

After request with `curl http://localhost:26657/broadcast_tx_commit?tx="1"`, `BroadcastTxCommit()` from https://github.com/cometbft/cometbft/blob/v0.37.x/rpc/core/mempool.go#L62 is executed. And then `CheckTx()` of abci-counter is invoked through socket connection. If tx successfully included in the block, `DeliverTx()` and `Commit()` of abci-counter is invoked.

