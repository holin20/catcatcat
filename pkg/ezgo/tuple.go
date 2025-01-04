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

type tuple3[T1, T2, T3 any] struct {
	_1 T1
	_2 T2
	_3 T3
}

func Pack3[T1, T2, T3 any](_1 T1, _2 T2, _3 T3) *tuple3[T1, T2, T3] {
	return &tuple3[T1, T2, T3]{_1, _2, _3}
}

func (p *tuple3[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return p._1, p._2, p._3
}
