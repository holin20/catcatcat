package ezgo

import (
	"fmt"
	"reflect"
	"strings"
)

func LoadFrom[T any](db *PostgresDB, v *T, st *structSqlTag[T]) error {
	sb, err := NewSqlBuilder().Select(st.fieldNames...).From(st.structName).Build()
	if IsErr(err) {
		return NewCause(err, "NewSqlBuilder")
	}
	colNames, rows, err := db.Query(sb)
	if IsErr(err) {
		return NewCause(err, "db.Query")
	}

	for _, r := range rows {
		for i, colVal := range r {
			colName := colNames[i]
			fieldName := st.tagToFieldName[colName]
			if fieldName == "" {
				return NewCausef(fmt.Errorf("no field name has tag: %s", colName), "tagToFieldName")
			}
			if err := setStructField(v, fieldName, colVal); IsErr(err) {
				return NewCause(err, "setStructField")
			}
		}
	}

	return nil
}

func Actualize[T any](db *PostgresDB, st *structSqlTag[T]) error {
	createTableLines := make([]string, len(st.fieldNames))
	for i, fieldName := range st.fieldNames {
		psqlType, err := goTypeKindToPostgresqlType(st.fieldNameToType[fieldName])
		if IsErr(err) {
			return NewCausef(err, "goTypeKindToPostgresqlType")
		}
		createTableLines[i] = "\t" + fieldName + " " + psqlType
	}

	createTableSql := fmt.Sprintf(
		`
CREATE TABLE %s (
	__id SERIAL8 PRIMARY KEY,
	%s
)`,
		st.structName,
		strings.Join(createTableLines, ",\n"),
	)

	_, err := db.db.Exec(createTableSql)
	if IsErr(err) {
		return NewCause(err, createTableSql)
	}

	return nil
}

func Create[T any](db *PostgresDB, v *T, st *structSqlTag[T]) error {
	sqlCols := make(map[string]*SqlCol)

	for _, fieldName := range st.fieldNames {
		field := reflect.ValueOf(v).Elem().FieldByName(fieldName)
		switch field.Kind() {
		case reflect.Bool:
			sqlCols[fieldName] = SqlColBool(field.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			sqlCols[fieldName] = SqlColInt(field.Int()) // TODO - make SqlColInt support int4
		case reflect.Int64:
			sqlCols[fieldName] = SqlColInt(field.Int())
		case reflect.Uint64:
			return fmt.Errorf("uint64 is not supported. Use int64 instead")
		case reflect.Float32, reflect.Float64:
			sqlCols[fieldName] = SqlColFloat(field.Float())
		case reflect.String:
			sqlCols[fieldName] = SqlColString(field.String())
		default:
			return fmt.Errorf("unsupported type: %s", field.Kind().String())
		}
	}

	return db.Insert(st.structName, sqlCols)
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
