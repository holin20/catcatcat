package orm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

const internalIdColName = "__id"
const internalTimeColName = "__ts" // unix time in milliseconds

// Load returns id->ent map
func Load[T any](db *ezgo.PostgresDB, schema *Schema[T], ids ...int64) (map[int64]*T, error) {
	// how to avoid allocating memory here?
	internalCols := []string{internalIdColName}
	colsToSelect := append(internalCols, schema.cols...)
	idConstraint := ""
	if len(ids) > 0 {
		idsInString := ezgo.SliceApply(ids, func(_ int, id int64) string {
			return strconv.FormatInt(id, 10)
		})
		idConstraint = internalIdColName + " in (" + strings.Join(idsInString, ",") + ")"
	}
	sb, err := ezgo.NewSqlBuilder().
		Select(colsToSelect...).
		From(schema.tableName).
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
			fieldName := schema.colToField[colName]
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

func getConstraint[T any](schema *Schema[T], v *T) (string, error) {
	if schema.partitionKeyCol == "" {
		return "", fmt.Errorf("no queryable partition field")
	}
	fieldName := schema.colToField[schema.partitionKeyCol]
	if fieldName == "" {
		return "", fmt.Errorf("col not found: %s", schema.partitionKeyCol)
	}

	field := reflect.ValueOf(v).Elem().FieldByName(fieldName)
	switch field.Kind() {
	case reflect.Bool:
		if field.Bool() {
			return schema.partitionKeyCol, nil
		} else {
			return "NOT " + schema.partitionKeyCol, nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Int64:
		return fmt.Sprintf("%s = %d", schema.partitionKeyCol, field.Int()), nil
	case reflect.Uint64:
		return "", fmt.Errorf("uint64 is not supported. Use int64 instead")
	case reflect.Float32, reflect.Float64:
		return "", fmt.Errorf("float can't be partition key")
	case reflect.String:
		return fmt.Sprintf("%s = '%s'", schema.partitionKeyCol, field.String()), nil
	default:
		return "", fmt.Errorf("unsupported type: %s", field.Kind().String())
	}
}

// LoadLastN queries the last N items by creation time sort by creation time in desc order
func LoadLastN[T any](
	db *ezgo.PostgresDB,
	schema *Schema[T],
	constraint *T,
	count int,
) ([]*ezgo.Pair_[int64, *T], error) {
	// how to avoid allocating memory here?
	internalCols := []string{internalTimeColName}
	colsToSelect := append(internalCols, schema.cols...)
	whereConstraint, err := getConstraint(schema, constraint)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "getConstraint")
	}
	sb, err := ezgo.NewSqlBuilder().
		Select(colsToSelect...).
		From(schema.tableName).
		Where(whereConstraint).
		OrderBy(map[string]ezgo.SqlOrderType{
			internalTimeColName: ezgo.SqlOrderDesc,
		}).
		Limit(count).
		Build()
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "NewSqlBuilder")
	}
	colNames, rows, err := db.Query(sb)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "db.Query")
	}

	var results []*ezgo.Pair_[int64, *T]
	for _, r := range rows {
		var ts int64 = -1
		var v T
		for ci, colVal := range r {
			colName := colNames[ci]
			if colName == internalTimeColName {
				ts = colVal.(int64)
				continue
			}
			fieldName := schema.colToField[colName]
			if fieldName == "" {
				return nil, fmt.Errorf("no field name has tag: %s", colName)
			}
			if err := setStructField(&v, fieldName, colVal); ezgo.IsErr(err) {
				return nil, ezgo.NewCause(err, "setStructField")
			}
		}
		if ts == -1 {
			return nil, fmt.Errorf("no ts found for this row")
		}
		results = append(results, ezgo.Pair(ts, &v))
	}

	return results, nil
}

func Actualize[T any](db *ezgo.PostgresDB, schema *Schema[T]) error {
	tableExisted, err := tableExists(db, schema.tableName)
	if ezgo.IsErr(err) {
		return ezgo.NewCausef(err, "tableExists: %s", schema.tableName)
	}
	if tableExisted {
		return alterSchema(db, schema)
	}

	createTableLines := make([]string, len(schema.fields))
	for i, fieldName := range schema.fields {
		psqlType, err := goTypeKindToPostgresqlType(schema.fieldProperty[fieldName].goType)
		if ezgo.IsErr(err) {
			return ezgo.NewCausef(err, "goTypeKindToPostgresqlType")
		}
		createTableLine := fmt.Sprintf(
			"\t%s %s",
			schema.fieldProperty[fieldName].sqlColName,
			psqlType,
		)
		if schema.fieldProperty[fieldName].isUinque {
			createTableLine += " UNIQUE"
		}
		createTableLines[i] = createTableLine
	}

	createTableSql := fmt.Sprintf(
		`
CREATE TABLE %s (
	%s SERIAL8 PRIMARY KEY,
	%s INT8,
%s
)`,
		schema.tableName,
		internalIdColName,
		internalTimeColName,
		strings.Join(createTableLines, ",\n"),
	)

	_, err = db.Exec(createTableSql)
	if ezgo.IsErr(err) {
		return ezgo.NewCause(err, createTableSql)
	}

	return nil
}

func alterSchema[T any](db *ezgo.PostgresDB, schema *Schema[T]) error {
	return fmt.Errorf("alterSchema unimplemented")
}

func tableExists(db *ezgo.PostgresDB, tableName string) (bool, error) {
	_, rows, err := db.Query(
		fmt.Sprintf(`SELECT count(*) FROM information_schema.tables where table_name = '%s'`, tableName),
	)
	if ezgo.IsErr(err) {
		return false, ezgo.NewCause(err, "db.Query")
	}
	return rows[0][0].(int64) == 1, nil
}

func Create[T any](db *ezgo.PostgresDB, schema *Schema[T], v *T) error {
	sqlCols := make(map[string]*ezgo.SqlCol)

	for fieldName, property := range schema.fieldProperty {
		// TODO - avoid using reflection to extract field value
		field := reflect.ValueOf(v).Elem().FieldByName(fieldName)
		switch field.Kind() {
		case reflect.Bool:
			sqlCols[property.sqlColName] = ezgo.SqlColBool(field.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Int64:
			sqlCols[property.sqlColName] = ezgo.SqlColInt(field.Int())
		case reflect.Uint64:
			return fmt.Errorf("uint64 is not supported. Use int64 instead")
		case reflect.Float32, reflect.Float64:
			sqlCols[property.sqlColName] = ezgo.SqlColFloat(field.Float())
		case reflect.String:
			sqlCols[property.sqlColName] = ezgo.SqlColString(field.String())
		default:
			return fmt.Errorf("unsupported type: %s", field.Kind().String())
		}
	}

	sqlCols[internalTimeColName] = ezgo.SqlColInt(time.Now().UnixMilli())

	return db.Insert(schema.tableName, sqlCols)
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
