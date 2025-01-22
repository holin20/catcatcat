package ezgo

import (
	"fmt"
	"strconv"
	"strings"
)

// Extremely simple sql builder that only supports basic operations

type SqlAggregateType string

const (
	AggregateCount SqlAggregateType = "COUNT"
	AggregateSum   SqlAggregateType = "SUM"
	AggregateAvg   SqlAggregateType = "AVG"
	AggregateMin   SqlAggregateType = "MIN"
	AggregateMax   SqlAggregateType = "MAX"
)

type PostgresqlNativeType string

const (
	PostgresqlTypeInt4   PostgresqlNativeType = "INT4"
	PostgresqlTypeInt8   PostgresqlNativeType = "INT8"
	PostgresqlTypeFloat4 PostgresqlNativeType = "FLOAT4"
	PostgresqlTypeFloat8 PostgresqlNativeType = "FLOAT8"
	PostgresqlTypeText   PostgresqlNativeType = "TEXT"
	PostgresqlTypeBool   PostgresqlNativeType = "BOOL"
)

//

type sqlColType uint8

const (
	sqlColTypeUnknown sqlColType = 0
	sqlColTypeString  sqlColType = 1
	sqlColTypeInt     sqlColType = 2
	sqlColTypeFloat   sqlColType = 3
	sqlColTypeBool    sqlColType = 4
)

type SqlCol struct {
	typ sqlColType
	str string
	i   int64
	f   float64
}

func SqlColInt(i int64) *SqlCol {
	return &SqlCol{
		typ: sqlColTypeInt,
		i:   i,
	}
}

func SqlColString(s string) *SqlCol {
	return &SqlCol{
		typ: sqlColTypeString,
		str: s,
	}
}

func SqlColFloat(f float64) *SqlCol {
	return &SqlCol{
		typ: sqlColTypeFloat,
		f:   f,
	}
}

func SqlColBool(b bool) *SqlCol {
	return &SqlCol{
		typ: sqlColTypeBool,
		i:   int64(If(b, 1, 0)),
	}
}

func (scs *SqlCol) String() string {
	switch scs.typ {
	case sqlColTypeString:
		return fmt.Sprintf("'%s'", scs.str)
	case sqlColTypeInt:
		return strconv.FormatInt(scs.i, 10)
	case sqlColTypeFloat:
		return fmt.Sprintf("%f", scs.f)
	case sqlColTypeBool:
		return If(scs.i == 1, "true", "false")
	}
	Fatalf("Unsupported sql col type: %d", scs.typ)
	return ""
}

//

type SqlBuilder struct {
	selectFields    Set[string]
	aggregateFields map[string]SqlAggregateType
	from            string
	constraint      string
	groupByFields   Set[string]
	orderByFields   Set[string]
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{
		selectFields:    MakeSet[string](),
		aggregateFields: make(map[string]SqlAggregateType),
		groupByFields:   MakeSet[string](),
		orderByFields:   MakeSet[string](),
	}
}

func (sb *SqlBuilder) Select(fields ...string) *SqlBuilder {
	sb.selectFields.Add(fields...)
	return sb
}

func (sb *SqlBuilder) Aggregate(field string, aggregate SqlAggregateType) *SqlBuilder {
	sb.aggregateFields[field] = aggregate
	return sb
}

func (sb *SqlBuilder) From(from string) *SqlBuilder {
	sb.from = from
	return sb
}

func (sb *SqlBuilder) Where(constraint string) *SqlBuilder {
	sb.constraint = constraint
	return sb
}

func (sb *SqlBuilder) GroupBy(fields ...string) *SqlBuilder {
	sb.groupByFields.Add(fields...)
	return sb
}

func (sb *SqlBuilder) OrderBy(fields ...string) *SqlBuilder {
	sb.orderByFields.Add(fields...)
	return sb
}

func (sb *SqlBuilder) Build() (string, error) {
	if sb.selectFields.Empty() {
		return "", fmt.Errorf("select fields empty")
	}

	// TODO - derive aggregate based on group-by fields
	selectClause := "SELECT " + strings.Join(sb.selectFields.ToSlice(), ", ")

	fromClause := If(sb.from != "", "FROM "+sb.from, "")
	whereClause := If(sb.constraint != "", "WHERE "+sb.constraint, "")

	// group by
	if !sb.groupByFields.CoveredBy(sb.selectFields) {
		return "", fmt.Errorf("some group-by fields do not exist in select: %s", sb.groupByFields.Substract(sb.selectFields).ToSlice())
	}
	groupByClause := If(!sb.groupByFields.Empty(), "GROUP BY "+strings.Join(sb.groupByFields.ToSlice(), ", "), "")

	// order by
	if !sb.orderByFields.CoveredBy(sb.selectFields) {
		return "", fmt.Errorf("some order-by fields do not exist in select: %s", sb.orderByFields.Substract(sb.selectFields).ToSlice())
	}
	orderByClause := If(!sb.orderByFields.Empty(), "ORDER BY "+strings.Join(sb.orderByFields.ToSlice(), ", "), "")

	clauses := []string{
		selectClause,
		fromClause,
		whereClause,
		groupByClause,
		orderByClause,
	}

	return strings.Join(SliceNonEmptyStringFilter(clauses), "\n"), nil
}

// INSERT INTO table_name (column1, column2, column3, ...)
// VALUES (value1, value2, value3, ...);

func BuildInsertSql(table string, cols map[string]*SqlCol) string {
	colNames := make([]string, len(cols))
	colValueStrings := make([]string, len(cols))
	i := 0
	for colName, colVal := range cols {
		colNames[i] = colName
		colValueStrings[i] = colVal.String()
		i++
	}
	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(colNames, ", "),
		strings.Join(colValueStrings, ", "),
	)
}
