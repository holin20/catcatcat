package monitor

import (
	"context"
	"time"
)

type Queryable[V any] interface {
	Query(ctx context.Context, now time.Time) (V, error)
}

// Echo

type Echo[V any] struct{ v V }

func NewEcho[V any](v V) *Echo[V] {
	return &Echo[V]{v: v}
}

func (e *Echo[V]) Query(ctx context.Context, now time.Time) (V, error) {
	return e.v, nil
}
