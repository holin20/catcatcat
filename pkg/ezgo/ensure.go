package ezgo

func EnsureMapValue[K comparable, V any](m map[K]*V, k K) *V {
	v := m[k]
	if v == nil {
		var zero V
		v = &zero
		m[k] = v
	}
	return v
}
