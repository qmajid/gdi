package main

import (
	"fmt"
	"test/di"
	"test/service"
)

type serviceA = service.A

func main() {
	//-------------Generic DI
	d := di.GNew()
	d.Provide(new(serviceA)).Provide(service.NewLogger())
	sa := di.Invoke[serviceA](d)
	sa.Print(1234)

	l := di.Invoke[service.ConcreteLogger](d)
	l.Info("hello di")

	//-------------Interface DI
	fmt.Print("------------ di with interface ------------ \n")
	d2 := di.INew()
	a := new(serviceA)
	l2 := service.NewLogger()
	d2.IProvide(a).IProvide(l2)
	
	si := d2.IInvoke(a.Name()).(*serviceA)
	si.Print(1024)
	li := d2.IInvoke(l2.Name()).(service.ILogger)
	li.Info("info message")
	li.Debug("debug message")
	li.Error("error message")
}
