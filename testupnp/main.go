package main

import (
	"fmt"
	"bytes"
	"net"
	"time"
)

func main() {
	ssdp, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if err != nil {
		return
	}
	packet, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		return
	}
	con := packet.(*net.UDPConn)

	defer con.Close()

	buf := bytes.NewBufferString(
		"M-SEARCH * HTTP/1.1\r\n" +
			"HOST: 239.255.255.250:1900\r\n" +
			"ST: urn:schemas-upnp-org:device:InternetGatewayDevice:1\r\n" +
			"MAN: \"ssdp:discover\"\r\n" +
			"MX: 2\r\n\r\n")

	message := buf.Bytes()

	err = con.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		return
	}

	_, err = con.WriteToUDP(message, ssdp)
	if err != nil {
		return
	}

	answerBytes := make([]byte, 1024)

	for{
		var n int
		n, _, err = con.ReadFromUDP(answerBytes)

		if err != nil {
			continue
		}

		if n <= 0 {
			continue
		}

		answer := string(answerBytes[0:n])

		fmt.Println(answer)
	}

}