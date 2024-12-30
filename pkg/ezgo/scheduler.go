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

type scheduler struct {
	scope *Scope
	wg    *sync.WaitGroup
}

func NewScheduler(scope *Scope) *scheduler {
	return &scheduler{
		scope: scope,
		wg:    &sync.WaitGroup{},
	}
}

func (s *scheduler) Repeat(
	ctx context.Context,
	interval time.Duration,
	task *task,
) {
	s.RepeatN(ctx, interval, -1, task)
}

func (s *scheduler) RepeatN(
	ctx context.Context,
	interval time.Duration,
	repeat int64, // negative number means infinite
	task *task,
) {
	ticker := time.NewTicker(interval)
	remaining := repeat
	s.wg.Add(1)
	go func() {
		defer ticker.Stop()
		defer s.wg.Done()
		for {
			if remaining == 0 {
				return
			}
			remaining--
			s.scope.GetLogger().Info(
				"Running periodic task",
				zap.String("task", task.GetName()),
				If(repeat < 0, zap.Skip(), zap.Int64("repeat", repeat)),
				If(repeat < 0, zap.Skip(), zap.Int64("remaining", remaining)),
			)
			task.Run()

			// block until next tick
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *scheduler) Once(
	ctx context.Context,
	after time.Duration,
	task *task,
) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		select {
		case <-time.After(after):
			task.Run()
		case <-ctx.Done():
			return
		}
	}()
}

func (s *scheduler) Join() {
	s.wg.Wait()
}
