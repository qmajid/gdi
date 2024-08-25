package main

import (
	"context"
	"fmt"
	"test/di"
	"test/service"
	"time"

	"C"
)
import (
	"log"
	"os"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

type Demo struct {
	A int
}

// func testMacro() {
// 	as := []Demo{}
// 	for i := 0; i < 1000; i++ {
// 		as = append(as, Demo{A: rand.Intn(100)})
// 	}

// 	bs := []Demo{}
// 	for i := range as {
// 		if as[i].A < 50 {
// 			bs = append(bs, as[i])
// 		}
// 	}
// 	fmt.Printf("bs=%d demos\n", len(bs))

// 	cs := FILTER(as, as[0].A < 50, Demo)

// 	fmt.Printf("cs=%d demos\n", len(cs))
// }

// #include "macros/functions.h"
func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// go worker(ctx)

	// time.Sleep(2 * time.Second)
	// cancel()
	// time.Sleep(1 * time.Second) // Give worker time to finish

	// testMacro()

	d := di.New()
	d.Provide(new(service.A)).Provide(log.New(os.Stdout, "majid ", 0))
	sa := di.Invoke[service.A](d)
	sa.Print(1234)

	// d.Provide(log.New(os.Stdout, "majid ", 0))
	l := di.Invoke[log.Logger](d)
	l.Printf("hello di")

	//-------------IDI
	fmt.Print("------------ di with interface ------------ \n")
	d2 := di.INew()
	a := new(service.A)
	l2 := service.NewLogger()
	d2.IProvide(a).IProvide(l2)
	si := d2.IInvoke(a.Name()).(*service.A)
	si.Print(1024)
	li := d2.IInvoke(l2.Name()).(*service.MyLogger)
	li.Print()
}
