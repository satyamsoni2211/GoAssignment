package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/satyamsoni2211/rpcbot/rpcserver"
)

func main() {
	api := new(rpcserver.API)
	rpc.Register(api)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Error encountered %s", e)
	}
	http.Serve(l, nil)
}
