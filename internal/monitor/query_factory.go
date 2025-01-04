package monitor

import (
	"encoding/csv"
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type QueryType int

const (
	FloatCsvQuery QueryType = 1
)

func BuildQuery[T float64](
	typ QueryType,
	args ...any,
) (Queryable[T], error) {
	switch typ {
	case FloatCsvQuery:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		csvReader := ezgo.AssertType[*csv.Reader](args[0], "FloatCsvQuery needs a csvReadr")
		q := &FloatCsvReaderQuery[T]{csvReader: *csvReader}
		return q, nil
	}
	return nil, fmt.Errorf("non-supported query type: %d", typ)
}
