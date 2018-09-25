package main

import (
    "encoding/json"
    "fmt"
    "net"
)

const (
    msg_length = 1024
)

func Echo(c net.Conn) {
    data := make([]byte, msg_length)
    defer c.Close()

    var recvdata map[string]string
    recvdata = make(map[string]string, 2)
    var senddata map[string]string
    senddata = make(map[string]string, 2)

    for {
        n, err := c.Read(data)
        if err != nil {
            fmt.Printf("read message from lotus failed")
            return
        }

        if err := json.Unmarshal(data[0:n], &recvdata); err == nil {
            senddata["reqId"] = recvdata["reqId"]
            senddata["resContent"] = "Hello " + recvdata["reqContent"]

            sendjson, err := json.Marshal(senddata)
            _, err = c.Write([]byte(sendjson))
            if err != nil {
                fmt.Printf("disconnect from lotus server")
                return
            }
        }
    }
}

func main() {
    fmt.Printf("Server is ready...\n")
    l, err := net.Listen("tcp", ":1988")
    if err != nil {
        fmt.Printf("Failure to listen: %s\n", err.Error())
    }

    for {
        if c, err := l.Accept(); err == nil {
            go Echo(c) //new thread
        }
    }
}