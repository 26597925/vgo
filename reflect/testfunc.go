package main

import (
	"fmt"
)

type Test struct {
	Name string
	Age int
}

func (this Test) Stringtest() {
	fmt.Println(this.Name)
}

func (this Test) Stringtest2() {
	fmt.Println("Stringtest2", this.Name)
}

func (this *Test) Stringtest1(i int) {
	this.Age = i
	fmt.Println(this.Age)
}

func GetNumArgs(args []string, num int) string {
	if num ==0 {
		num = 1
	}
	return args[num]
}