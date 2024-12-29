package gen

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
