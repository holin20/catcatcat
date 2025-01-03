package monitor

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type Rule[V any] struct {
	query     Queryable[V]
	condition Condition[V]
}

func NewRule[V any]() *Rule[V] {
	return &Rule[V]{}
}

func (r *Rule[V]) Eval(ctx context.Context, now time.Time) (bool, error) {
	queryResult, err := r.query.Query(ctx, now)
	if ezgo.IsErr(err) {
		return false, ezgo.NewCause(err, "Query")
	}
	return r.condition.Met(queryResult), nil
}
