package main

import (
	"fmt"
)

func decorator(f func(s string)) func(s string) {

	return func (s string) {
			fmt.Println("Started")
			f(s)
			fmt.Println("End")
	}
}

func hello(s string) {
	fmt.Println(s)
}

func main() {
	decorator(hello)("hello world")

	h := decorator(hello)
	h("Hello")
}