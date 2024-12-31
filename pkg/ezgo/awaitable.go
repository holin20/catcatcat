package ezgo

import "sync"

type Awaitable struct {
	doneChan chan struct{}
}

func NewAwaitable() (*Awaitable, func()) {
	doneChan := make(chan struct{})
	awaitable := &Awaitable{
		doneChan: doneChan,
	}
	signalFunc := func() {
		close(doneChan)
	}

	return awaitable, signalFunc
}

func (a *Awaitable) Await() {
	<-a.doneChan
}

func Await(awaitables ...Awaitable) {
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

//////

type Awaitable1[T any] struct {
	doneChan chan T
}

func NewAwaitable1[T any]() (*Awaitable1[T], func(v T)) {
	doneChan := make(chan T)
	awaitable := &Awaitable1[T]{
		doneChan: doneChan,
	}
	signalFunc := func(v T) {
		doneChan <- v
	}

	return awaitable, signalFunc
}

func (a *Awaitable1[T]) Await() T {
	return <-a.doneChan
}

func Await1[T any](awaitables ...Awaitable1[T]) []T {
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
