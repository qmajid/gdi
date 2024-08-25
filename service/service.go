package service

import (
	"fmt"
)

type serviceA interface {
	Print(n int)
}

type A struct {
	Number int
}

func (a A) Print(n int) {
	fmt.Printf("A-> %+v\n", n)
}

// -----------------
type B struct {
	Name string
}

func (b B) Print() {
	fmt.Printf("B-> %+v\n", b)
}
