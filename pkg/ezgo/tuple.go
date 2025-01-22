package ezgo

type Tuple2_[T1, T2 any] struct {
	_1 T1
	_2 T2
}

func Tuple2[T1, T2 any](_1 T1, _2 T2) *Tuple2_[T1, T2] {
	return &Tuple2_[T1, T2]{_1, _2}
}

func (t *Tuple2_[T1, T2]) Unpack() (T1, T2) {
	return t._1, t._2
}

type Pair_[T1, T2 any] Tuple2_[T1, T2]

func Pair[T1, T2 any](_1 T1, _2 T2) *Pair_[T1, T2] {
	return &Pair_[T1, T2]{_1, _2}
}

func (t *Pair_[T1, T2]) Unpack() (T1, T2) {
	return t._1, t._2
}

type Tuple3_[T1, T2, T3 any] struct {
	_1 T1
	_2 T2
	_3 T3
}

func Pack3[T1, T2, T3 any](_1 T1, _2 T2, _3 T3) *Tuple3_[T1, T2, T3] {
	return &Tuple3_[T1, T2, T3]{_1, _2, _3}
}

func (p *Tuple3_[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return p._1, p._2, p._3
}
