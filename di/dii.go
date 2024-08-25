package di

import (
	"fmt"
	"test/service"
)

type iservice service.Service

type idiRegistery struct {
	m map[string]iservice
}

func INew() *idiRegistery {
	s := new(idiRegistery)
	s.m = make(map[string]iservice)
	return s
}

func (s *idiRegistery) IProvide(inst iservice) *idiRegistery {
	name := inst.Name()
	s.m[inst.Name()] = inst
	fmt.Printf("di Provide name:[%v]\n", name)
	return s
}

func (s *idiRegistery) IInvoke(name string) iservice {
	fmt.Printf("di Invoke name:[%v]\n", name)
	return s.m[name]
	// reflect.ValueOf(inst).MethodByName("Print").Call([]reflect.Value{})
	// return &inst
}
