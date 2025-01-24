package ezgo

import (
	"reflect"
)

type StructTag[T any] struct {
	FieldNameToTag  map[string]string
	TagToFieldName  map[string]string
	FieldNameToType map[string]reflect.Kind
	FieldTags       []string
	FieldNames      []string
	StructName      string
}

func NewStructTag[T any](tag string) *StructTag[T] {
	var v T
	t := reflect.TypeOf(v)
	Assertf(t.Kind() == reflect.Struct, "type %s is not struct", t.Name())

	st := StructTag[T]{
		FieldNameToTag:  make(map[string]string),
		FieldNameToType: make(map[string]reflect.Kind),
		TagToFieldName:  make(map[string]string),
		StructName:      t.Name(),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldTag := string(field.Tag.Get(tag))

		if fieldTag == "" {
			continue // warn on empty tag?
		}

		st.FieldNameToTag[field.Name] = fieldTag
		st.FieldNameToType[field.Name] = field.Type.Kind()

		Assertf(st.TagToFieldName[fieldTag] == "", "duplicated tag: %s", fieldTag)
		st.TagToFieldName[string(field.Tag.Get(tag))] = field.Name

		st.FieldTags = append(st.FieldTags, fieldTag)
		st.FieldNames = append(st.FieldNames, field.Name)
	}

	return &st
}
