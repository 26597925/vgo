package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"errors"
	"strings"
)

func chooseListenPort(nat NAT, externalPort int) (listenPort int, err error) {
	// TODO: Unmap port when exiting. (Right now we never exit cleanly.)
	// TODO: Defend the port, remap when router reboots
	listenPort, err = nat.AddPortMapping("tcp", externalPort, externalPort,
		"Taipei-Torrent port "+strconv.Itoa(externalPort), 360000)
	if err != nil {
		return
	}
	return
}

func GetLocalIntenetIP() string {
	/*
	  获得所有本机地址
	  判断能联网的ip地址
	*/

	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		panic(errors.New("不能连接网络"))
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Println("Start Upnp")

		fmt.Println(GetLocalIntenetIP())

		nat, err := Discover()

		log.Println("Upnp loading")
		if err != nil {
			err = fmt.Errorf("Unable to create NAT: %v", err)
			log.Println(err)
			return
		}

		log.Println("Upnp loading")

		listenPort := 8191
		if nat != nil {
			var external net.IP
			if external, err = nat.GetExternalAddress(); err != nil {
				err = fmt.Errorf("Unable to get external IP address from NAT: %v", err)
				return
			}
			log.Println("External ip address: ", external)

			if listenPort, err = chooseListenPort(nat, listenPort); err != nil {
				log.Println("Could not choose listen port.", err)
				log.Println("Peer connectivity will be affected.")
			}
		}

		log.Println("Upnp end")

		wg.Done()
	}()

	wg.Wait()
}
