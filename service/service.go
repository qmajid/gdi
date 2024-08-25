package service

import (
	"fmt"
)

type Service interface {
	Do()
}

type A struct {
	Number int
}

func (a A) Do() {
	fmt.Printf("call Do in A -> %+v\n", a)
}

func (a A) Print(n int) {
	fmt.Printf("A::Print -> %v\n", n)
}

// -----------------
type B struct {
	Name string
}

func (b B) Print() {
	fmt.Printf("B-> %+v\n", b)
}

func (b B) Do() {
	fmt.Printf("call Do in B -> %+v\n", b)
}
