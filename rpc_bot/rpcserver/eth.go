package rpcserver

import (
	"fmt"

	"github.com/umbracle/go-web3"
)

type API struct {
}

func (api *API) Greet(greet string, reply *string) error {
	fmt.Println("Received Message ", greet)
	*reply = `Please enter your blockchain address 
	(eg. 0xE4559721E46326F18FF59e3D926E4489FE6a5162)`
	return nil
}

func (api *API) ValidateEthAdress(adress string, reply *string) error {
	ethAdress := web3.Address{}
	err := ethAdress.UnmarshalText([]byte(adress))
	if err != nil {
		fmt.Println(err)
		*reply = "Invalid Blockchain Address"
		return nil
	}
	*reply = fmt.Sprintf("Received Blockchain address %s", ethAdress.String())
	return nil
}
