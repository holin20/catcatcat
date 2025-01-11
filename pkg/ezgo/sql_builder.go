package ezgo

import (
	"fmt"
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

type SqlBuilder struct {
	selectFields    Set[string]
	aggregateFields map[string]SqlAggregateType
	from            string
	constraint      string
	groupByFields   Set[string]
	orderByFields   Set[string]
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
