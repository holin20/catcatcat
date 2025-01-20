package ezgo

import "reflect"

type structTag[T any] struct {
	tagCache map[string]string
}

func StructTag[T any](tag string) structTag[T] {
	var v T
	t := reflect.TypeOf(v)
	Assertf(t.Kind() == reflect.Struct, "type %s is not struct", t.Name())

	st := structTag[T]{tagCache: make(map[string]string)}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		st.tagCache[field.Name] = string(field.Tag.Get(tag))
	}

	return st
}

func (st *structTag[T]) GetField(field string) string {
	return st.tagCache[field]
}
