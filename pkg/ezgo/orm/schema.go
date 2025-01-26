package orm

import (
	"reflect"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const sqlColTag = "sql"

type Schema[T any] struct {
	fieldToCol    map[string]string
	colToField    map[string]string
	fieldToGoType map[string]reflect.Kind
	cols          []string
	fields        []string
	tableName     string
}

func NewSchema[T any]() *Schema[T] {
	var v T
	t := reflect.TypeOf(v)
	ezgo.Assertf(t.Kind() == reflect.Struct, "type %s is not struct", t.Name())

	schema := Schema[T]{
		fieldToCol:    make(map[string]string),
		fieldToGoType: make(map[string]reflect.Kind),
		colToField:    make(map[string]string),
		tableName:     ezgo.CamelToSnake(t.Name()),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := string(field.Tag.Get(sqlColTag))

		if col == "" {
			continue // warn on empty tag?
		}

		schema.fieldToCol[field.Name] = col
		schema.fieldToGoType[field.Name] = field.Type.Kind()

		ezgo.Assertf(schema.colToField[col] == "", "duplicated col: %s", col)
		schema.colToField[string(field.Tag.Get(sqlColTag))] = field.Name

		schema.cols = append(schema.cols, col)
		schema.fields = append(schema.fields, field.Name)
	}

	return &schema
}
