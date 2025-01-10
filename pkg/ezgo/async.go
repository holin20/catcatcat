package ezgo

func AsyncVoid(fn func()) *AwaitableVoid {
	awaitable, signal := NewAwaitableVoid()
	go func() {
		defer signal()
		fn()
	}()
	return awaitable
}

func Async[T any](fn func() T) *Awaitable[T] {
	awaitable, signal := NewAwaitable[T]()
	go func() {
		signal(fn())
	}()
	return awaitable
}

func Async2[T1, T2 any](fn func() (T1, T2)) *Awaitable[*Tuple2_[T1, T2]] {
	awaitable, signal := NewAwaitable[*Tuple2_[T1, T2]]()
	go func() {
		signal(Tuple2(fn()))
	}()
	return awaitable
}
