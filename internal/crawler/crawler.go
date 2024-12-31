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
	items := []struct {
		name     string
		url      string
		interval time.Duration
	}{
		{
			name: "Macbook Air M3 15\" 16G 256GB",
			url:  "https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620577642&catalogId=10701&productId=3074457345620577640&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
	}

	for _, item := range items {
		item := item
		interval := item.interval
		if interval == 0 {
			interval = 10 * time.Second
		}
		c.scheduler.Repeat(ctx, interval, ezgo.NewNamedTask("Fetch "+item.name, func() {
			price, err := costco.FetchItemPrice(item.url)
			if ezgo.IsErr(err) {
				ezgo.LogCauses(c.scope.GetLogger(), err, "FetchPrice")
				return
			}
			c.scope.GetLogger().Info("Fetched Item Info", zap.String("item", item.name), zap.Float64("price", price))
		}))
	}

	// c.scheduler.Repeat(ctx, 10*time.Second, ezgo.NewNamedTask("FetchMacbookPrice", func() {
	// 	price, err := costco.FetchMacbookPrice()
	// 	if ezgo.IsErr(err) {
	// 		ezgo.LogCauses(c.scope.GetLogger(), err, "FetchMacbookPrice")
	// 		return
	// 	}
	// 	c.scope.GetLogger().Info("Price", zap.Float64("price", price))
	// }))
}

func (c *Crawler) Terminate() *ezgo.Awaitable {
	return c.scheduler.Terminate()
}
