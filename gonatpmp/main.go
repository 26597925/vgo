package main

import (
	"fmt"
    "github.com/jackpal/gateway"
    natpmp "github.com/jackpal/go-nat-pmp"
)

func main() {
	gatewayIP, err := gateway.DiscoverGateway()
	fmt.Println(gatewayIP)

	if err != nil {
		return
	}
	
	fmt.Println("==========================1")
	client := natpmp.NewClient(gatewayIP)
	fmt.Println("==========================2")
	response, err := client.GetExternalAddress()
	fmt.Println("==========================3")

	if err != nil {
		return
	}
	fmt.Println("==========================4")

	ip := string(response.ExternalIPAddress[:])
	fmt.Println("External IP address:"+ ip)

	for {

	}
}