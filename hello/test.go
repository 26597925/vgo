package main

import (
	"fmt"
)

type hello struct {
	world string
}
func main() {
	hlo := hello{}
	hlo.world = "test"
	fmt.Printf("%s", hlo.world)
}