package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func connect() (client *rpc.Client, err error) {
	client, err = rpc.DialHTTP("tcp", "localhost:1234")
	return
}

func main() {
	client, err := connect()
	if err != nil {
		log.Fatalf("Unable to connect to rpc server, is server running ? err: %s", err)
	}
	reader := bufio.NewReader(os.Stdin)
	var response, api string
	fmt.Println("< Hi there from RPC bot, you can greet me or ask me to validate eth adress, /q to quit")
	for {
		fmt.Print("> ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("< Error reding string, please retype ... ")
			continue
		}
		msg = strings.Trim(msg, "\n")
		toLowerMsg := strings.ToLower(msg)
		switch toLowerMsg {
		case "hi", "hello", "hey", "hola", "hi there":
			api = "API.Greet"
		case "quit", "/q", "q", "exit":
			fmt.Println("Thanks for using rpc bot, bbye")
			return
		default:
			api = "API.ValidateEthAdress"
		}
		err = client.Call(api, msg, &response)
		if err != nil {
			fmt.Printf("< unrecognized response from RPC  server, %s\n", err)
			continue
		}
		fmt.Printf("< %s \n", response)
	}
}
