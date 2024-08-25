package main

import (
	"fmt"
	"test/di"
	"test/service"
)

func main() {
	//-------------Generic DI
	d := di.GNew()
	d.Provide(new(service.A)).Provide(service.NewLogger())
	sa := di.Invoke[service.A](d)
	sa.Print(1234)

	l := di.Invoke[service.MyLogger](d)
	l.Info("hello di")

	//-------------Interface DI
	fmt.Print("------------ di with interface ------------ \n")
	d2 := di.INew()
	a := new(service.A)
	l2 := service.NewLogger()
	d2.IProvide(a).IProvide(l2)
	si := d2.IInvoke(a.Name()).(*service.A)
	si.Print(1024)
	li := d2.IInvoke(l2.Name()).(service.ILogger)
	li.Info("info message")
	li.Debug("debug message")
	li.Error("error message")
}
