package monitor

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type Rule[V any] struct {
	name      string
	query     Queryable[V]
	condition Condition[V]
}

func NewRule[V any]() *Rule[V] {
	return &Rule[V]{}
}

func (r *Rule[V]) Eval(ctx context.Context, now time.Time) (bool, V, error) {
	queryResult, err := r.query.Query(ctx, now)
	var zero V
	if ezgo.IsErr(err) {
		return false, zero, ezgo.NewCause(err, "Query")
	}
	return r.condition.Met(queryResult), queryResult, nil
}

func (r *Rule[V]) GetName() string {
	return r.name
}
