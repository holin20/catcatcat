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
		doneChan <- struct{}{}
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
