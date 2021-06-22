package main

import (
	"flag"
	"fmt"
	"github.com/semrush/zenrpc/v2"
	"log"
	"net/http"
	"os"
)

type MetaMask struct{ zenrpc.Service }

func (as MetaMask) Eth_blockNumber() int {
	return 5
}
func (as MetaMask) Net_version() int {
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
	fmt.Println(*addr)

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true, AllowCORS: true})
	rpc.Register("", MetaMask{}) // public
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting arithsrv on %s", "localhost:8545")
	server := &http.Server{Addr: ":8545", Handler: nil}
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("failed to start metamask service")
		}


}
