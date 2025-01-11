package ezgo

type Set[K comparable] struct {
	m map[K]struct{}
}

func MakeSet[K comparable](keys ...K) Set[K] {
	s := Set[K]{
		m: make(map[K]struct{}),
	}
	for _, k := range keys {
		s.m[k] = struct{}{}
	}
	return s
}

func (s Set[K]) Has(k K) bool {
	_, ok := s.m[k]
	return ok
}

func (s Set[K]) Add(keys ...K) {
	for _, k := range keys {
		s.m[k] = struct{}{}
	}
}

func (s Set[K]) Remove(keys ...K) {
	for _, k := range keys {
		delete(s.m, k)
	}
}

func (s Set[K]) Size() int {
	return len(s.m)
}

func (s Set[K]) Empty() bool {
	return s.Size() == 0
}

func (s Set[K]) ToSlice() []K {
	ret := make([]K, s.Size())
	offset := 0
	for k := range s.m {
		ret[offset] = k
		offset++
	}
	return ret
}

func (s Set[K]) Covers(t Set[K]) bool {
	for tk := range t.m {
		if !s.Has(tk) {
			return false
		}
	}
	return true
}

func (s Set[K]) CoveredBy(t Set[K]) bool {
	return t.Covers(s)
}

func (s Set[K]) Substract(t Set[K]) Set[K] {
	sub := MakeSet[K]()
	for sk := range s.m {
		if !t.Has(sk) {
			sub.Add(sk)
		}
	}
	return sub
}
