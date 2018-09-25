package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{}
	proc, err := os.StartProcess("master vlock", []string{}, attr)
	fmt.Println(err)
	fmt.Println(proc)
}