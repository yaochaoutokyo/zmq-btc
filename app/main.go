package main

import (
	"fmt"
	"github.com/btcsuite/btcutil"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	subZmq("tcp://3.112.178.254:28332")
}

func subZmq(uri string) {
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if nil != err {
		panic(err)
	}
	defer subscriber.Close()
	err = subscriber.Connect(uri)
	if nil != err {
		panic(err)
	}
	err = subscriber.SetSubscribe("rawtx")
	if nil != err {
		panic(err)
	}
	err = subscriber.SetRcvtimeo(300 * time.Minute)
	if nil != err {
		panic(err)
	}
	for {
		_, err := subscriber.Recv(0)
		if nil != err {
			panic(err)
		}
		txByte, err := subscriber.RecvBytes(0)
		if nil != err {
			panic(err)
		}
		_, _ = subscriber.Recv(0)
		tx, err := btcutil.NewTxFromBytes(txByte)
		if err != nil {
			panic(err)
		}
		fmt.Println("txhash ", tx.Hash().String())
	}
}
