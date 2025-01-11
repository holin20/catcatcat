package ezgo

import "sync"

type AwaitableVoid struct {
	wg *sync.WaitGroup
}

func NewAwaitableVoid() (*AwaitableVoid, func()) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	awaitable := &AwaitableVoid{wg: wg}
	signalFunc := func() { wg.Done() }
	return awaitable, signalFunc
}

func (a *AwaitableVoid) Await() {
	a.wg.Wait()
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
	wg *sync.WaitGroup
	v  T
}

func NewAwaitable[T any]() (*Awaitable[T], func(v T)) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	awaitable := &Awaitable[T]{wg: wg}
	signalFunc := func(v T) {
		awaitable.v = v
		wg.Done()
	}
	return awaitable, signalFunc
}

func (a *Awaitable[T]) Await() T {
	a.wg.Wait()
	return a.v
}

func AwaitAll[T any](awaitables ...*Awaitable[T]) []T {
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

func AwaitMapAll[K comparable, T any](awaitables map[K]*Awaitable[T]) map[K]T {
	flattenedResult := make([]T, len(awaitables))
	seqToKey := make(map[int]K)
	wg := sync.WaitGroup{}
	var i int
	for k, a := range awaitables {
		seqToKey[i] = k
		wg.Add(1)
		go func(seq int) {
			defer wg.Done()
			flattenedResult[seq] = a.Await()
		}(i)
		i++
	}
	wg.Wait()
	resultMap := make(map[K]T)
	for i, r := range flattenedResult {
		resultMap[seqToKey[i]] = r
	}
	return resultMap
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
