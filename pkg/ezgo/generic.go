package ezgo

func Must[T any](v T, err error) T {
	AssertNoError(err, "Must")
	return v
}

func In[T comparable](item T, items ...T) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

func IfLazy[T any](cond bool, trueFunc, falseFunc func() T) T {
	if cond {
		return trueFunc()
	}
	return falseFunc()
}

func IfEval[T any](cond bool, trueFunc, falseFunc func()) {
	if cond {
		trueFunc()
	} else {
		falseFunc()
	}
}

func NonEmptyOr(x, y string) string {
	return If(x != "", x, y)
}

func NonNilOr[T any](x, y *T) *T {
	return If(x != nil, x, y)
}

func Arg1[T1, Trest any](arg1 T1, args ...Trest) T1 {
	return arg1
}

func Arg2[T1, T2, Trest any](arg1 T1, arg2 T2, args ...Trest) T2 {
	return arg2
}

func Arg3[T1, T2, T3, Trest any](arg1 T1, arg2 T2, arg3 T3, args ...Trest) T3 {
	return arg3
}
