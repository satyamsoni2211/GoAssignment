package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	client.Call("API.Hello", "satyam", &name)
	fmt.Println(name)
}
