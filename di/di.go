package di

import (
	"fmt"
	"reflect"
)

type diRegistery struct {
	m map[string]interface{}
}

func New() *diRegistery {
	s := new(diRegistery)
	s.m = make(map[string]interface{})
	return s
}

func structName(i any) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (s *diRegistery) Provide(inst any) *diRegistery {
	name := structName(inst)
	s.m[name] = inst
	fmt.Printf("di Provide name:[%v]\n", name)
	return s
}

func Invoke[T any](s *diRegistery) *T {
	var inst T
	name := structName(inst)
	fmt.Printf("di Invoke name:[%v]\n", name)
	return s.m[name].(*T)
	// reflect.ValueOf(inst).MethodByName("Print").Call([]reflect.Value{})
	// return &inst
}
