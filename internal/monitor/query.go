package monitor

import (
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/holin20/catcatcat/pkg/ezgo/orm"
)

var (
	cdpSchema = orm.NewSchema[schema.Cdp]()
)

type Queryable[V any] interface {
	Query(ctx context.Context, now time.Time) (V, time.Time, error)
}

// FloatCsvReaderQuery

type FloatCsvReaderQuery[T float64] struct{ csvReader csv.Reader }

func NewFloatCsvReaderQuery[T float64](csvReader csv.Reader) *FloatCsvReaderQuery[T] {
	return &FloatCsvReaderQuery[T]{csvReader: csvReader}
}

func (e *FloatCsvReaderQuery[T]) Query(ctx context.Context, now time.Time) (T, time.Time, error) {
	var zero T
	cols, err := e.csvReader.Read()
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCause(err, "readCsvPart")
	}
	if len(cols) != 1 {
		return zero, now, ezgo.NewCause(fmt.Errorf("one row should only have 1 col but got %d", len(cols)), "assert_col_cnt")
	}
	v, err := strconv.ParseFloat(cols[0], 64)
	ezgo.IsOk(err)
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCausef(err, "ParseFloat(%s)", cols[0])
	}

	return T(v), now, nil
}

// ZapTailQuery

type ZapTailQuery[T float64] struct {
	logPath   string
	timeField string
	valField  string
}

func NewZapTailQuery[T float64](logPath, timeField, valField string) *ZapTailQuery[T] {
	return &ZapTailQuery[T]{logPath, timeField, valField}
}

func (e *ZapTailQuery[T]) Query(ctx context.Context, now time.Time) (T, time.Time, error) {
	lastLine, err := ezgo.TailFile(e.logPath)
	var zero T
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCause(err, "TailFile")
	}
	if len(lastLine) == 0 { // empty file
		return zero, now, ezgo.NewCause(fmt.Errorf("empty file: %s", e.logPath), "EmptyFile")
	}
	result, err := ezgo.ExtractJsonPath(lastLine, e.valField)
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s) for value", lastLine, e.valField)
	}

	ts, err := ezgo.ExtractJsonPath(lastLine, e.timeField)
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s) for time", lastLine, e.timeField)
	}

	parsedTime, err := time.Parse(time.RFC3339, ts.Str)
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCausef(err, "time.Parse(%s)", ts.Str)
	}

	return T(result.Float()), parsedTime, nil
}

// PostgresSqlQuery

type PostgresSqlQuery[T float64] struct {
	db        *ezgo.PostgresDB
	sqlString string
	timeField string
	valField  string
}

func NewPostgresSqlQuery[T float64](
	db *ezgo.PostgresDB,
	sqlString string,
	timeField string,
	valField string,
) *PostgresSqlQuery[T] {
	return &PostgresSqlQuery[T]{db, sqlString, timeField, valField}
}

func (p *PostgresSqlQuery[T]) Query(ctx context.Context, now time.Time) (T, time.Time, error) {
	colNames, rows, err := p.db.Query(p.sqlString)
	if ezgo.IsErr(err) {
		return 0, time.UnixMicro(0), ezgo.NewCause(err, "PostgresSqlQuery_Query")
	}
	timeColIndex, valColIndex := -1, -1
	for i, name := range colNames {
		switch name {
		case p.timeField:
			timeColIndex = i
		case p.valField:
			valColIndex = i
		}
	}
	if timeColIndex == -1 || valColIndex == -1 {
		return 0, time.UnixMicro(0), ezgo.NewCause(fmt.Errorf("timeColIndex/valColIndex not found"), "PostgresSqlQuery_locate_time_val_index")
	}
	val := rows[0][valColIndex].(float64)
	ts := rows[0][timeColIndex].(int64)

	return T(val), time.Unix(ts, 0), nil
}

// PostgresSqlQuery

type EntCdpQuery[T schema.Cdp] struct {
	db    *ezgo.PostgresDB
	catId string
}

func NewEntCdpQuery[T schema.Cdp](
	db *ezgo.PostgresDB,
	catId string,
) *EntCdpQuery[T] {
	return &EntCdpQuery[T]{db: db, catId: catId}
}

func (q *EntCdpQuery[T]) Query(ctx context.Context, now time.Time) (T, time.Time, error) {
	results, err := orm.LoadLastN(q.db, cdpSchema, &schema.Cdp{CatId: q.catId}, 1)
	var zero T
	if ezgo.IsErr(err) {
		return zero, time.UnixMicro(0), ezgo.NewCause(err, "LoadLastN")
	}
	if len(results) == 0 {
		return zero, time.UnixMicro(0), fmt.Errorf("zero result from LoadLastN")
	}
	ts, cdp := results[0].Unpack()
	return T(*cdp), time.UnixMilli(ts), nil
}
