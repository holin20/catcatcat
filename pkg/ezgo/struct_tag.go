package ezgo

import (
	"reflect"
)

type structTag[T any] struct {
	fieldNameToTag  map[string]string
	tagToFieldName  map[string]string
	fieldNameToType map[string]reflect.Kind
	offsetCache     map[string]uintptr
	fieldNames      []string
	structName      string
}

func StructTag[T any](tag string) structTag[T] {
	var v T
	t := reflect.TypeOf(v)
	Assertf(t.Kind() == reflect.Struct, "type %s is not struct", t.Name())

	st := structTag[T]{
		fieldNameToTag: make(map[string]string),
		fieldNames:     make([]string, t.NumField()),
		structName:     t.Name(),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldTag := string(field.Tag.Get(tag))

		st.fieldNameToTag[field.Name] = fieldTag
		st.fieldNameToType[field.Name] = field.Type.Kind()

		Assertf(st.tagToFieldName[fieldTag] == "", "duplicated tag: %s", fieldTag)
		st.tagToFieldName[string(field.Tag.Get(tag))] = field.Name

		st.offsetCache[field.Name] = field.Offset
		st.fieldNames[i] = field.Name
	}

	return st
}

func (st *structTag[T]) GetFieldNameForTag(tag string) string {
	return st.tagToFieldName[tag]
}

func (st *structTag[T]) GetTagForFieldName(fieldName string) string {
	return st.fieldNameToTag[fieldName]
}

func (st *structTag[T]) GetFieldNames() []string {
	return st.fieldNames
}

func (st *structTag[T]) GetOffset(field string) uintptr {
	return st.offsetCache[field]
}

//

type structSqlTag[T any] structTag[T]

func StructSqlTag[T any]() structTag[T] {
	return StructTag[T]("sql")
}
