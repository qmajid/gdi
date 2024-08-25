package service

import (
	"fmt"
	"log"
	"os"
)

type Service interface {
	Name() string
}

type A struct {
	Number int
}

func (a A) Name() string {
	return "A"
}

func (a A) Print(n int) {
	fmt.Printf("A::Print -> %v\n", n)
}

// -----------------
type B struct {
	// name string
}

func (b B) Print() {
	fmt.Printf("B-> %+v\n", b)
}

func (b B) Name() string {
	return "B"
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

func (l MyLogger) Name() string {
	return "MyLogger"
}
