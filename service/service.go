package service

import (
	"fmt"
	"log"
	"os"
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

//----------------

type MyLogger struct {
	*log.Logger
}

func New() *MyLogger {
	m := MyLogger{log.New(os.Stdout, "majid ", 0)}
	// m.l = log.New(os.Stdout, "majid ", 0)
	return &m
}

func (l MyLogger) Do() {
	fmt.Printf("call Do in MyLogger\n")
}
