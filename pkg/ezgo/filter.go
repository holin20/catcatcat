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

// Common filters

func SliceFilterTrue(slice []bool) []bool {
	return SliceFilter(slice, func(v bool) bool { return v })
}

func SliceFilterFalse(slice []bool) []bool {
	return SliceFilter(slice, func(v bool) bool { return !v })
}

func SliceFilterNonEmptyString(slice []string) []string {
	return SliceFilter(slice, func(s string) bool { return s != "" })

}
