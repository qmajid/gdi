package di

import (
	"test/service"
)

type iservice service.Service

type idi struct {
	m map[string]iservice
}

func INew() *idi {
	s := new(idi)
	s.m = make(map[string]iservice)
	return s
}

func (s *idi) IProvide(inst iservice) *idi {
	name := inst.Name()
	s.m[name] = inst
	// fmt.Printf("di Provide name:[%v]\n", name)
	return s
}

func (s *idi) IInvoke(name string) iservice {
	// fmt.Printf("di Invoke name:[%v]\n", name)
	return s.m[name]
	// reflect.ValueOf(inst).MethodByName("Print").Call([]reflect.Value{})
	// return &inst
}
