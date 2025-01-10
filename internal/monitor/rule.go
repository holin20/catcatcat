package monitor

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type Rule[V any] struct {
	id        string
	name      string
	query     Queryable[V]
	condition Condition[V]
}

func (r *Rule[V]) Eval(ctx context.Context, now time.Time) (bool, V, time.Time, error) {
	queryResult, qtime, err := r.query.Query(ctx, now)
	var zero V
	if ezgo.IsErr(err) {
		return false, zero, now, ezgo.NewCause(err, "Query")
	}
	return r.condition.Met(queryResult), queryResult, qtime, nil
}

func (r *Rule[V]) GetName() string {
	return r.name
}
