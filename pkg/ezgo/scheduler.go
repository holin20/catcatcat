package ezgo

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Task

type task struct {
	name string
	fn   func()
}

func (t *task) GetName() string {
	if t.name != "" {
		return t.name
	}
	return "[unnamed]"
}

func (t *task) Run() {
	t.fn()
}

func NewNamedTask(name string, fn func()) *task {
	return &task{name: name, fn: fn}
}

func NewUnnamedTask(fn func()) *task {
	return &task{fn: fn}
}

// Scheduler

type Scheduler struct {
	scope    *Scope
	wg       *sync.WaitGroup
	doneChan chan struct{}
}

func NewScheduler(scope *Scope) *Scheduler {
	return &Scheduler{
		scope:    scope.WithLogger(scope.GetLogger().Named("Scheduler")),
		wg:       &sync.WaitGroup{},
		doneChan: make(chan struct{}),
	}
}

func (s *Scheduler) Repeat(
	ctx context.Context,
	interval time.Duration,
	task *task,
) {
	s.RepeatN(ctx, interval, -1, task)
}

func (s *Scheduler) RepeatN(
	ctx context.Context,
	interval time.Duration,
	repeat int64, // negative number means infinite
	task *task,
) {
	ticker := time.NewTicker(interval)
	remaining := repeat
	seq := 0
	s.wg.Add(1)
	go func() {
		defer ticker.Stop()
		defer s.wg.Done()
		for {
			if remaining == 0 {
				return
			}
			remaining--
			seq++
			s.scope.GetLogger().Info(
				"Running periodic task",
				zap.String("task", task.GetName()),
				zap.Int("seq", seq),
				If(repeat < 0, zap.Skip(), zap.Int64("repeat", repeat)),
				If(repeat < 0, zap.Skip(), zap.Int64("remaining", remaining)),
			)
			task.Run()

			// block until next tick
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				s.scope.GetLogger().Debug("Received ctx.Done() in RepeatN")
				return
			case <-s.doneChan:
				s.scope.GetLogger().Debug("Received doneChan in RepeatN")
				return
			}
		}
	}()
}

func (s *Scheduler) Once(
	ctx context.Context,
	after time.Duration,
	task *task,
) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		select {
		case <-time.After(after):
			s.scope.GetLogger().Info("Running one-off task", zap.String("task", task.GetName()))
			task.Run()
		case <-ctx.Done():
			s.scope.GetLogger().Debug("Received ctx.Done() in Once")
			return
		case <-s.doneChan:
			s.scope.GetLogger().Debug("Received doneChan in Once")
			return
		}
	}()
}

func (s *Scheduler) Terminate() *Awaitable {
	awaitable, signal := NewAwaitable()
	go func() {
		defer signal()
		s.wg.Wait()
	}()

	s.doneChan <- struct{}{}

	return awaitable
}
