package ezgo

import "sync"

type AwaitableVoid struct {
	doneChan chan struct{}
}

func NewAwaitableVoid() (*AwaitableVoid, func()) {
	doneChan := make(chan struct{})
	awaitable := &AwaitableVoid{
		doneChan: doneChan,
	}
	signalFunc := func() {
		close(doneChan)
	}

	return awaitable, signalFunc
}

func (a *AwaitableVoid) Await() {
	<-a.doneChan
}

func AwaitAvoid(awaitables ...AwaitableVoid) {
	wg := sync.WaitGroup{}
	for _, a := range awaitables {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Await()
		}()
	}
	wg.Wait()
}

////////

type Awaitable[T any] struct {
	doneChan chan T
}

func NewAwaitable[T any]() (*Awaitable[T], func(v T)) {
	doneChan := make(chan T)
	awaitable := &Awaitable[T]{
		doneChan: doneChan,
	}
	signalFunc := func(v T) {
		doneChan <- v
	}

	return awaitable, signalFunc
}

func (a *Awaitable[T]) Await() T {
	return <-a.doneChan
}

func Await[T any](awaitables ...*Awaitable[T]) []T {
	result := make([]T, len(awaitables))
	wg := sync.WaitGroup{}
	for i, a := range awaitables {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = a.Await()
		}()
	}
	wg.Wait()
	return result
}

func Await2[T1, T2 any](awaitable1 *Awaitable[T1], awaitable2 *Awaitable[T2]) (T1, T2) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	var r1 T1
	go func() {
		defer wg.Done()
		r1 = awaitable1.Await()
	}()

	wg.Add(1)
	var r2 T2
	go func() {
		defer wg.Done()
		r2 = awaitable2.Await()
	}()

	wg.Wait()
	return r1, r2
}

func Await3[T1, T2, T3 any](a1 *Awaitable[T1], a2 *Awaitable[T2], a3 *Awaitable[T3]) (T1, T2, T3) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	var r1 T1
	go func() {
		defer wg.Done()
		r1 = a1.Await()
	}()

	wg.Add(1)
	var r2 T2
	go func() {
		defer wg.Done()
		r2 = a2.Await()
	}()

	wg.Add(1)
	var r3 T3
	go func() {
		defer wg.Done()
		r3 = a3.Await()
	}()

	wg.Wait()

	return r1, r2, r3
}
