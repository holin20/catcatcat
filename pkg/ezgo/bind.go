package ezgo

// Bind 1 arg

func Bind1[I1 any](fn func(I1), arg1 I1) func() {
	return func() {
		fn(arg1)
	}
}

func Bind1_1[I1, O1 any](fn func(I1) O1, arg1 I1) func() O1 {
	return func() O1 {
		return fn(arg1)
	}
}

func Bind1_2[I1, O1 any](fn func(I1) O1, arg1 I1) func() O1 {
	return func() O1 {
		return fn(arg1)
	}
}

// Bind 2 args

func Bind2[I1, I2 any](fn func(I1, I2), arg1 I1, arg2 I2) func() {
	return func() {
		fn(arg1, arg2)
	}
}

func Bind2_1[I1, I2, O1 any](fn func(I1, I2) O1, arg1 I1, arg2 I2) func() O1 {
	return func() O1 {
		return fn(arg1, arg2)
	}
}

func Bind2_2[I1, I2, O1, O2 any](fn func(I1, I2) (O1, O2), arg1 I1, arg2 I2) func() (O1, O2) {
	return func() (O1, O2) {
		return fn(arg1, arg2)
	}
}

// Bind 3 args

func Bind3[I1, I2, I3 any](fn func(I1, I2, I3), arg1 I1, arg2 I2, arg3 I3) func() {
	return func() {
		fn(arg1, arg2, arg3)
	}
}

func Bind3_1[I1, I2, I3, O1 any](fn func(I1, I2, I3) O1, arg1 I1, arg2 I2, arg3 I3) func() O1 {
	return func() O1 {
		return fn(arg1, arg2, arg3)
	}
}

func Bind3_2[I1, I2, I3, O1, O2 any](fn func(I1, I2, I3) (O1, O2), arg1 I1, arg2 I2, arg3 I3) func() (O1, O2) {
	return func() (O1, O2) {
		return fn(arg1, arg2, arg3)
	}
}

// Bind 4 args

func Bind4[I1, I2, I3, I4 any](fn func(I1, I2, I3, I4), arg1 I1, arg2 I2, arg3 I3, arg4 I4) func() {
	return func() {
		fn(arg1, arg2, arg3, arg4)
	}
}

func Bind4_1[I1, I2, I3, I4, O1 any](fn func(I1, I2, I3, I4) O1, arg1 I1, arg2 I2, arg3 I3, arg4 I4) func() O1 {
	return func() O1 {
		return fn(arg1, arg2, arg3, arg4)
	}
}

func Bind4_2[I1, I2, I3, I4, O1, O2 any](fn func(I1, I2, I3, I4) (O1, O2), arg1 I1, arg2 I2, arg3 I3, arg4 I4) func() (O1, O2) {
	return func() (O1, O2) {
		return fn(arg1, arg2, arg3, arg4)
	}
}
