package main

import (
	"github.com/pixelbender/go-stun/stun"
	"fmt"
)

func main() {
    conn, addr, err := stun.Discover("stun:stun1.l.google.com:19302")
	if err != nil {
    	fmt.Println(err)
    	return
    }
    defer conn.Close()
	fmt.Printf("Local address: %v, Server reflexive address: %v", conn.LocalAddr(), addr)
}