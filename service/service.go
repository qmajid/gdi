package service

import (
	"fmt"
	"log"
	"os"
)

// --------------------- INTERFACE
type Service interface {
	Name() string
}

// --------------------- CONCRETE STRUCTS
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

type ILogger interface {
	Service
	Info(s string)
	Debug(s string)
	Error(s string)
}

type MyLogger struct {
	l *log.Logger
}

func (l MyLogger) Name() string {
	return "MyLogger"
}

func (l MyLogger) Info(s string) {
	l.l.Printf("[I]:pring from logger [%v]\n", s)
}

func (l MyLogger) Debug(s string) {
	l.l.Printf("[D]:pring from logger [%v]\n", s)
}

func (l MyLogger) Error(s string) {
	l.l.Printf("[E]:pring from logger [%v]\n", s)
}

func NewLogger() ILogger {
	m := MyLogger{l: log.New(os.Stdout, "majid: ", 0)}
	return &m
}
