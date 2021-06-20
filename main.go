package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/semrush/zenrpc/v2"
)

type MetaMask struct{ zenrpc.Service }

func (as MetaMask) Eth_blockNumber() int {
	fmt.Println("gotcha")
	return 5
}
func (as MetaMask) Net_version() int {
	fmt.Println("gotcha")
	return 3
}

func (as MetaMask) Eth_getBalance(addr string, blockNumber int) int {
	return 100
}

func (as MetaMask) Eth_getBlockByNumber(blockNumber int, filter bool) string {
	return "blockContent"
}






//go:generate zenrpc

func main() {
	addr := flag.String("addr", "localhost:8545", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("", MetaMask{}) // public
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting arithsrv on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
