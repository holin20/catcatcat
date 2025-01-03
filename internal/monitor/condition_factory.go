package monitor

import (
	"cmp"
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type ConditionType int

const (
	eGreater ConditionType = 1
	eLess    ConditionType = 2
	eEqual   ConditionType = 3
	eInside  ConditionType = 4
	eOutside ConditionType = 5
)

func BuildCondition[T cmp.Ordered](
	typ ConditionType,
	args ...T,
) (Condition[T], error) {
	switch typ {
	case eGreater:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Greater[T]{Threshold: args[0]}, nil
	case eLess:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Less[T]{Threshold: args[0]}, nil
	case eEqual:
		ezgo.Assertf(len(args) == 1, "len(args) should be euqal to %d for type %d", 1, typ)
		return &Equal[T]{Target: args[0]}, nil
	}
	return nil, fmt.Errorf("non-supported condition type: %d", typ)
}
