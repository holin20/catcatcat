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
