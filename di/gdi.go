package di

import (
	"reflect"
)

type gdi struct {
	m map[string]interface{}
}

func GNew() *gdi {
	s := new(gdi)
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

func (s *gdi) Provide(inst any) *gdi {
	name := structName(inst)
	s.m[name] = inst
	// fmt.Printf("di Provide name:[%v]\n", name)
	return s
}

func Invoke[T any](s *gdi) *T {
	var inst T
	name := structName(inst)
	// fmt.Printf("di Invoke name:[%v]\n", name)
	return s.m[name].(*T)
	// reflect.ValueOf(inst).MethodByName("Print").Call([]reflect.Value{})
	// return &inst
}
