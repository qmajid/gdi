package main

import (
	"fmt"
	"log"
	"os"
	"test/di"
	"test/service"
)

func main() {

	d := di.GNew()
	d.Provide(new(service.A)).Provide(log.New(os.Stdout, "majid ", 0))
	sa := di.Invoke[service.A](d)
	sa.Print(1234)

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
