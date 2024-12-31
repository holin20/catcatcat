package ezgo

func FlattenMap[K comparable, V1 any, V2 any](m map[K]V1, flattenFunc func(key K, value V1) V2) []V2 {
	flattened := make([]V2, len(m))
	i := 0
	for k, v := range m {
		flattened[i] = flattenFunc(k, v)
		i++
	}
	return flattened
}
