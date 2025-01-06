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
	Query(ctx context.Context, now time.Time) (V, error)
}

// FloatCsvReaderQuery

type FloatCsvReaderQuery[T float64] struct{ csvReader csv.Reader }

func NewFloatCsvReaderQuery[T float64](csvReader csv.Reader) *FloatCsvReaderQuery[T] {
	return &FloatCsvReaderQuery[T]{csvReader: csvReader}
}

func (e *FloatCsvReaderQuery[T]) Query(ctx context.Context, now time.Time) (T, error) {
	var zero T
	cols, err := e.csvReader.Read()
	if ezgo.IsErr(err) {
		return zero, ezgo.NewCause(err, "readCsvPart")
	}
	if len(cols) != 1 {
		return zero, ezgo.NewCause(fmt.Errorf("one row should only have 1 col but got %d", len(cols)), "assert_col_cnt")
	}
	v, err := strconv.ParseFloat(cols[0], 64)
	ezgo.IsOk(err)
	if ezgo.IsErr(err) {
		return zero, ezgo.NewCausef(err, "ParseFloat(%s)", cols[0])
	}

	return T(v), nil
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

func (e *ZapTailQuery[T]) Query(ctx context.Context, now time.Time) (T, error) {
	lastLine, err := ezgo.TailFile(e.logPath)
	var zero T
	if ezgo.IsErr(err) {
		return zero, ezgo.NewCause(err, "TailFile")
	}
	if len(lastLine) == 0 { // empty file
		return zero, ezgo.NewCause(fmt.Errorf("empty file: %s", e.logPath), "EmptyFile")
	}
	result, err := ezgo.ExtractJsonPath(lastLine, e.valField)
	if ezgo.IsErr(err) {
		return zero, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s)", lastLine, e.valField)
	}
	return T(result.Float()), nil
}
