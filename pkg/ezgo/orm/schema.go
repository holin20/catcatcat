package orm

import (
	"reflect"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const sqlColTag = "sql"
const isUniqueTag = "unique"
const partitionKeyTag = "pk"

type fieldProperty struct {
	sqlColName string
	isUinque   bool
	goType     reflect.Kind
}

type Schema[T any] struct {
	fieldProperty map[string]fieldProperty

	colToField map[string]string
	cols       []string // sql cols
	fields     []string // golang struct fields
	tableName  string

	partitionKeyCol string
}

func NewSchema[T any]() *Schema[T] {
	var v T
	t := reflect.TypeOf(v)
	ezgo.Assertf(t.Kind() == reflect.Struct, "type %s is not struct", t.Name())

	schema := Schema[T]{
		fieldProperty: make(map[string]fieldProperty),
		colToField:    make(map[string]string),
		tableName:     ezgo.CamelToSnake(t.Name()),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := string(field.Tag.Get(sqlColTag))
		isUnqiue := field.Tag.Get(isUniqueTag) == "true"
		isPartitionKey := field.Tag.Get(partitionKeyTag) == "true"

		if col == "" {
			continue // warn on empty tag?
		}

		if isPartitionKey {
			ezgo.Assertf(schema.partitionKeyCol == "", "Duplicated partition key: %s (previous %s)", col, schema.partitionKeyCol)
			schema.partitionKeyCol = col
		}

		schema.fieldProperty[field.Name] = fieldProperty{
			sqlColName: col,
			isUinque:   isUnqiue,
			goType:     field.Type.Kind(),
		}

		ezgo.Assertf(schema.colToField[col] == "", "duplicated col: %s", col)
		schema.colToField[string(field.Tag.Get(sqlColTag))] = field.Name

		schema.cols = append(schema.cols, col)
		schema.fields = append(schema.fields, field.Name)
	}

	return &schema
}
