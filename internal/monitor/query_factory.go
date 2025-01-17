package monitor

import (
	"encoding/csv"
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type QueryType int

const (
	FloatCsvReader QueryType = 1
	ZapTail        QueryType = 2
	PostgreSql     QueryType = 3
)

func BuildQuery[T float64](
	typ QueryType,
	args ...any,
) (Queryable[T], error) {
	switch typ {
	case FloatCsvReader:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		csvReader := ezgo.AssertType[*csv.Reader](args[0], "FloatCsvReader needs a csvReadr")
		q := &FloatCsvReaderQuery[T]{csvReader: *csvReader}
		return q, nil
	case ZapTail:
		ezgo.Assertf(len(args) == 3, "len(args) should be euqal to %d for type %d", 3, typ)
		logFilePath := ezgo.AssertType[string](args[0], "arg0 should be string")
		timeField := ezgo.AssertType[string](args[1], "arg1 should be string")
		valField := ezgo.AssertType[string](args[2], "arg2 should be string")
		q := NewZapTailQuery[T](logFilePath, timeField, valField)
		return q, nil
	case PostgreSql:
		ezgo.Assertf(len(args) == 4, "len(args) should be euqal to %d for type %d", 3, typ)
		db := ezgo.AssertType[*ezgo.PostgresDB](args[0], "arg0 should be PostgresDB")
		sqlString := ezgo.AssertType[string](args[1], "arg1 should be sqlString")
		timeField := ezgo.AssertType[string](args[2], "arg2 should be string")
		valField := ezgo.AssertType[string](args[3], "arg3 should be string")
		q := NewPostgresSqlQuery[T](db, sqlString, timeField, valField)
		return q, nil
	}
	return nil, fmt.Errorf("non-supported query type: %d", typ)
}
