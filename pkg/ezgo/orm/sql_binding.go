package ezgo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const internalIdColName = "__id"

func Load[T any](db *ezgo.PostgresDB, st *ezgo.StructTag[T], ids ...int64) (map[int64]*T, error) {
	// how to avoid allocating memory here?
	colsToSelect := append([]string{internalIdColName}, st.FieldTags...)
	idConstraint := ""
	if len(ids) > 0 {
		idsInString := ezgo.SliceApply(ids, func(_ int, id int64) string {
			return strconv.FormatInt(id, 10)
		})
		idConstraint = internalIdColName + " in (" + strings.Join(idsInString, ",") + ")"
	}
	sb, err := ezgo.NewSqlBuilder().
		Select(colsToSelect...).
		From(st.StructName).
		Where(idConstraint).
		Build()
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "NewSqlBuilder")
	}
	colNames, rows, err := db.Query(sb)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "db.Query")
	}

	results := make(map[int64]*T, len(rows))
	for _, r := range rows {
		var id int64 = -1
		var v T
		for ci, colVal := range r {
			colName := colNames[ci]
			if colName == internalIdColName {
				id = colVal.(int64)
				continue
			}
			fieldName := st.TagToFieldName[colName]
			if fieldName == "" {
				return nil, fmt.Errorf("no field name has tag: %s", colName)
			}
			if err := setStructField(&v, fieldName, colVal); ezgo.IsErr(err) {
				return nil, ezgo.NewCause(err, "setStructField")
			}
		}
		if id == -1 {
			return nil, fmt.Errorf("no id found for this row")
		}
		results[id] = &v
	}

	return results, nil
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
	%s SERIAL8 PRIMARY KEY,
%s
)`,
		st.StructName,
		internalIdColName,
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
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16:
		return "INT4", nil
	case reflect.Int, reflect.Uint32, reflect.Int64:
		return "INT8", nil
	case reflect.Uint, reflect.Uint64:
		return "", fmt.Errorf("Uint/Uint64 is not supported. Use int64 instead")
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
	if !valReflect.Type().ConvertibleTo(field.Type()) {
		return fmt.Errorf(
			"cannot covert v(%v. type: %s) to field %s (type: %s)",
			v,
			valReflect.Type().String(),
			fieldName,
			field.Type().String(),
		)
	}

	// if !valReflect.Type().AssignableTo(field.Type()) {
	// 	return fmt.Errorf(
	// 		"cannot assing v(%v. type: %s) to field %s (type: %s)",
	// 		v,
	// 		valReflect.Type().String(),
	// 		fieldName,
	// 		field.Type().String(),
	// 	)
	// }

	field.Set(valReflect.Convert(field.Type()))
	return nil
}
