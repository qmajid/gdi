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

//----------------

type MyLogger struct {
	l *log.Logger
}

func NewLogger() *MyLogger {
	m := MyLogger{l: log.New(os.Stdout, "majid: ", 0)}
	return &m
}

func (l MyLogger) Name() string {
	return "MyLogger"
}

func (l MyLogger) Print() {
	l.l.Printf("pring from logger\n")
}
