package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	abciserver "github.com/cometbft/cometbft/abci/server"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
)

type CounterApp struct {
	abci.BaseApplication
	counter int64
}

func (app *CounterApp) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	txVal, err := strconv.ParseInt(string(req.Tx), 10, 64)
	if err != nil || txVal != app.counter + 1 {
		return abci.ResponseCheckTx{Code: 1, Log: "Invalid tx"}
	}
	return abci.ResponseCheckTx{Code: 0}
}

func (app *CounterApp) DeliverTx(req abci.RequestDeliverTx) abci.ResponseDeliverTx {
	txVal, err := strconv.ParseInt(string(req.Tx), 10, 64)
	if err != nil || txVal != app.counter + 1 {
		return abci.ResponseDeliverTx{Code: 1, Log: "Invalid tx"}
	}
	app.counter = txVal
	return abci.ResponseDeliverTx{Code: 0}
}

func (app *CounterApp) Commit() abci.ResponseCommit {
	hash := make([]byte, 8)
	binary.BigEndian.PutUint64(hash, uint64(app.counter))
	return abci.ResponseCommit{Data: hash}
}

func main() {
	app := &CounterApp{}
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	srv, err := abciserver.NewServer("tcp://0.0.0.0:26658", "socket", app)
	if err != nil {
		panic(err)
	}
	srv.SetLogger(logger)

	if err := srv.Start(); err != nil {
		panic(err)
	}
	fmt.Println("ABCI Counter App running on port 26658...")

	select{}
}
