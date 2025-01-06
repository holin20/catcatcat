package monitor

import (
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
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

	parsedTime, err := time.Parse("2006-01-02 15:04:05", ts.Str)
	if ezgo.IsErr(err) {
		return zero, now, ezgo.NewCausef(err, "time.Parse(%s)", ts.Str)
	}

	return T(result.Float()), parsedTime, nil
}
