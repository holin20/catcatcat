package ezgo

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

func LoadFrom[T any](db *ezgo.PostgresDB, v *T, st *ezgo.StructTag[T]) error {
	sb, err := ezgo.NewSqlBuilder().Select(st.FieldNames...).From(st.StructName).Build()
	if ezgo.IsErr(err) {
		return ezgo.NewCause(err, "NewSqlBuilder")
	}
	colNames, rows, err := db.Query(sb)
	if ezgo.IsErr(err) {
		return ezgo.NewCause(err, "db.Query")
	}

	for _, r := range rows {
		for i, colVal := range r {
			colName := colNames[i]
			fieldName := st.TagToFieldName[colName]
			if fieldName == "" {
				return ezgo.NewCausef(fmt.Errorf("no field name has tag: %s", colName), "tagToFieldName")
			}
			if err := setStructField(v, fieldName, colVal); ezgo.IsErr(err) {
				return ezgo.NewCause(err, "setStructField")
			}
		}
	}

	return nil
}

func Actualize[T any](db *ezgo.PostgresDB, st *ezgo.StructTag[T]) error {
	createTableLines := make([]string, len(st.FieldNames))
	for i, fieldName := range st.FieldNames {
		psqlType, err := goTypeKindToPostgresqlType(st.FieldNameToType[fieldName])
		if ezgo.IsErr(err) {
			return ezgo.NewCausef(err, "goTypeKindToPostgresqlType")
		}
		createTableLines[i] = "\t" + st.FieldNameToTag[fieldName] + " " + psqlType
	}

	createTableSql := fmt.Sprintf(
		`
CREATE TABLE %s (
	__id SERIAL8 PRIMARY KEY,
%s
)`,
		st.StructName,
		strings.Join(createTableLines, ",\n"),
	)

	_, err := db.Exec(createTableSql)
	if ezgo.IsErr(err) {
		return ezgo.NewCause(err, createTableSql)
	}

	return nil
}

func Create[T any](db *ezgo.PostgresDB, v *T, st *ezgo.StructTag[T]) error {
	sqlCols := make(map[string]*ezgo.SqlCol)

	for fieldName, tag := range st.FieldNameToTag {
		// TODO - avoid using reflection to extract field value
		field := reflect.ValueOf(v).Elem().FieldByName(fieldName)
		switch field.Kind() {
		case reflect.Bool:
			sqlCols[tag] = ezgo.SqlColBool(field.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Int64:
			sqlCols[tag] = ezgo.SqlColInt(field.Int())
		case reflect.Uint64:
			return fmt.Errorf("uint64 is not supported. Use int64 instead")
		case reflect.Float32, reflect.Float64:
			sqlCols[tag] = ezgo.SqlColFloat(field.Float())
		case reflect.String:
			sqlCols[tag] = ezgo.SqlColString(field.String())
		default:
			return fmt.Errorf("unsupported type: %s", field.Kind().String())
		}
	}

	return db.Insert(st.StructName, sqlCols)
}

func goTypeKindToPostgresqlType(kind reflect.Kind) (string, error) {
	switch kind {
	case reflect.Bool:
		return "BOOL", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "INT4", nil
	case reflect.Int64:
		return "INT8", nil
	case reflect.Uint64:
		return "", fmt.Errorf("Uint64 is not supported. Use int64 instead")
	case reflect.Float32:
		return "FLOAT4", nil
	case reflect.Float64:
		return "FLOAT8", nil
	case reflect.String:
		return "TEXT", nil
	default:
		return "", fmt.Errorf("unsupported golang type: %s", kind.String())
	}
}

// setStructField sets field value for objPtr. Currently use reflection.
// TODO - use cached struct refelction to avoid runtime reflection.
func setStructField[T any](objPtr *T, fieldName string, v any) error {
	field := reflect.ValueOf(objPtr).Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("invalid field of struct: %s", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("field of struct cannot set: %s", fieldName)
	}
	valReflect := reflect.ValueOf(v)
	if !valReflect.Type().AssignableTo(field.Type()) {
		return fmt.Errorf("cannot assing v to field %s", fieldName)
	}

	field.Set(valReflect)
	return nil
}
