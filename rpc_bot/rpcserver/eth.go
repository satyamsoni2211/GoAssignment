package rpcserver

import "fmt"

type API struct {
}

func (api *API) Hello(name string, reply *string) error {
	fmt.Println("Received ", name)
	*reply = fmt.Sprintf("Hello %s", name)
	return nil
}
