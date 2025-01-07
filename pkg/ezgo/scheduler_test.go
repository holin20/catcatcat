package ezgo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduler(t *testing.T) {
	scheduler := NewScheduler(NewScopeWithDefaultLogger("TestScheduler"))

	ctx, cancel := context.WithCancel(context.Background())

	scheduler.RepeatN(ctx, 1*time.Second, 5, "", func() {
		fmt.Printf("%d\n", time.Now().UnixMilli())
	})

	time.Sleep(5 * time.Second)
	cancel()

	assert.Fail(t, "TestScheduler")
}
