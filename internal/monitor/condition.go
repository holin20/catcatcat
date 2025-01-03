package monitor

import "cmp"

type Condition[T any] interface {
	Met(val T) bool
}

// Greater

type Greater[T cmp.Ordered] struct{ Threshold T }

func (g *Greater[T]) Met(val T) bool { return val > g.Threshold }

// Less

type Less[T cmp.Ordered] struct{ Threshold T }

func (l *Less[T]) Met(val T) bool { return val < l.Threshold }

// Equal

type Equal[T cmp.Ordered] struct{ Val, Target T }

func (e *Equal[T]) Met(val T) bool { return val == e.Target }
