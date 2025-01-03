package monitor

import (
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type QueryType int

const (
	eEcho QueryType = 1
)

func BuildQuery[T any](
	typ QueryType,
	args ...any,
) (Queryable[T], error) {
	switch typ {
	case eEcho:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Echo[T]{}, nil
	}
	return nil, fmt.Errorf("non-supported query type: %d", typ)
}
