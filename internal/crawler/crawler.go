package crawler

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

type Crawler struct {
	scheduler *ezgo.Scheduler
	scope     *ezgo.Scope
}

func NewCrawler(scope *ezgo.Scope) *Crawler {
	scope = scope.WithLogger(scope.GetLogger().Named("Crawler"))
	return &Crawler{
		scheduler: ezgo.NewScheduler(scope),
		scope:     scope,
	}
}

func (c *Crawler) Start(ctx context.Context) {
	c.scheduler.Repeat(ctx, 10*time.Second, ezgo.NewNamedTask("FetchMacbookPrice", func() {
		price, err := costco.FetchMacbookPrice()
		if ezgo.IsErr(err) {
			ezgo.LogCauses(c.scope.GetLogger(), err, "FetchMacbookPrice")
			return
		}
		c.scope.GetLogger().Info("Price", zap.Float64("price", price))
	}))
}

func (c *Crawler) Terminate() *ezgo.Awaitable {
	return c.scheduler.Terminate()
}
