package aa

import (
	"fmt"
)

type Aa struct  {
	aa string
}

type Bb struct  {
	bb string
}

func (this Aa) GetAa() {
	fmt.Println(this.aa)
}

func (this Bb) getBb() {
	fmt.Println(this.bb)
}