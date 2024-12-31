package ezgo

func Async(fn func()) *Awaitable {
	awaitable, signal := NewAwaitable()
	go func() {
		defer signal()
		fn()
	}()
	return awaitable
}

func Async1[T any](fn func() T) *Awaitable1[T] {
	awaitable, signal := NewAwaitable1[T]()
	go func() {
		signal(fn())
	}()
	return awaitable
}
