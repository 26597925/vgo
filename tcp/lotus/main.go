package main

import (
    "lotuslib"
)

const (
    ip   = "0.0.0.0"
    port = 1987
)

func main() {
    tcplotus.TcpLotusMain(ip, port)
}