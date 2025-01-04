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

func SliceApply[I, O any](input []I, apply func(I) O) []O {
	result := make([]O, len(input))
	for i, x := range input {
		result[i] = apply(x)
	}
	return result
}

func SliceApplyAsync[I, O any](input []I, apply func(int, I) O) []*Awaitable[O] {
	awaitables := make([]*Awaitable[O], len(input))
	for i, x := range input {
		awaitables[i] = Async(Bind2_1(apply, i, x))
	}
	return awaitables
}
