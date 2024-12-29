package gen

func First[T1, Trest any](arg1 T1, args ...Trest) T1 {
	return arg1
}

func Second[T1, T2, Trest any](arg1 T1, arg2 T2, args ...Trest) T2 {
	return arg2
}
