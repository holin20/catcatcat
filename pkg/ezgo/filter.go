package ezgo

func SliceFilter[T any](slice []T, fn func(T) bool) []T {
	var ret []T
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func MapFilter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	ret := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			ret[k] = v
		}
	}
	return ret
}

// Common filters

func SliceTrueFilter(slice []bool) []bool {
	return SliceFilter(slice, func(v bool) bool { return v })
}

func SliceFalseFilter(slice []bool) []bool {
	return SliceFilter(slice, func(v bool) bool { return !v })
}

func SliceNonEmptyStringFilter(slice []string) []string {
	return SliceFilter(slice, func(s string) bool { return s != "" })

}

func MapTrueFilter[K comparable](m map[K]bool) map[K]bool {
	return MapFilter(m, func(k K, v bool) bool { return v })
}
