package ezgo

type tuple2[T1, T2 any] struct {
	_1 T1
	_2 T2
}

func Tuple2[T1, T2 any](_1 T1, _2 T2) *tuple2[T1, T2] {
	return &tuple2[T1, T2]{_1, _2}
}

func (t *tuple2[T1, T2]) Unpack() (T1, T2) {
	return t._1, t._2
}

type pack3[T1, T2, T3 any] struct {
	arg1 T1
	arg2 T2
	arg3 T3
}

func Pack3[T1, T2, T3 any](arg1 T1, arg2 T2, arg3 T3) *pack3[T1, T2, T3] {
	return &pack3[T1, T2, T3]{arg1, arg2, arg3}
}

func (p *pack3[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return p.arg1, p.arg2, p.arg3
}

// Getter

func Arg1[T1, Trest any](arg1 T1, args ...Trest) T1 {
	return arg1
}

func Arg2[T1, T2, Trest any](arg1 T1, arg2 T2, args ...Trest) T2 {
	return arg2
}

func Arg3[T1, T2, T3, Trest any](arg1 T1, arg2 T2, arg3 T3, args ...Trest) T3 {
	return arg3
}
