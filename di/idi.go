package di

import (
	"test/service"
)

// dependency injection by interface

type iservice service.Service

type idi struct {
	m map[string]iservice
}

func INew() *idi {
	return &idi{m: make(map[string]iservice)}
}

func (s *idi) IProvide(inst iservice) *idi {
	name := inst.Name()
	s.m[name] = inst
	return s
}

func (s *idi) IInvoke(name string) iservice {
	dep, ok := s.m[name]
	if !ok {
		return nil
	}
	return dep

	// reflect.ValueOf(inst).MethodByName("Print").Call([]reflect.Value{})
}
