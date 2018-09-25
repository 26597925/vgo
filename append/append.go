package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	const MaxUint = 1<<16 - 1
	fmt.Println(MaxUint)
}