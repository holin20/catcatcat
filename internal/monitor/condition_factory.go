package monitor

import (
	"cmp"
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type ConditionType int

const (
	GreaterCondition  ConditionType = 1
	LessCondition     ConditionType = 2
	EqualCondition    ConditionType = 3
	InsideCondition   ConditionType = 4
	eOutsideCondition ConditionType = 5
)

func BuildCondition[T cmp.Ordered](
	typ ConditionType,
	args ...T,
) (Condition[T], error) {
	switch typ {
	case GreaterCondition:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Greater[T]{Threshold: args[0]}, nil
	case LessCondition:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Less[T]{Threshold: args[0]}, nil
	case EqualCondition:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Equal[T]{Target: args[0]}, nil
	}
	return nil, fmt.Errorf("non-supported condition type: %d", typ)
}
