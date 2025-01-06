package crawler

import (
	"context"
	"fmt"
	"time"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/internal/model"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

const (
	defaultCrawlInterval = 1 * time.Minute
)

type CrawlListEntry = struct {
	CatId  string
	Cat    *model.Cat
	Costco *model.CostcoFetcher
}

type Crawler struct {
	scheduler    *ezgo.Scheduler
	scope        *ezgo.Scope
	crawlList    []CrawlListEntry
	crawInterval time.Duration
}

func NewCrawler(scope *ezgo.Scope) *Crawler {
	scope = scope.WithLogger(scope.GetLogger().Named("Crawler"))
	return &Crawler{
		scheduler: ezgo.NewScheduler(scope),
		scope:     scope,
	}
}

func (c *Crawler) WithCrawlList(crawlList []CrawlListEntry) *Crawler {
	c.crawlList = crawlList
	return c
}

func (c *Crawler) WithCrawlInterval(crawlInterval time.Duration) *Crawler {
	c.crawInterval = crawlInterval
	return c
}

func (c *Crawler) Kickoff(ctx context.Context) {
	for _, entry := range c.crawlList {
		entry := entry
		interval := ezgo.NonZeroOr(c.crawInterval, defaultCrawlInterval)
		resultLogger := ezgo.CloneLogger(
			c.scope.GetLogger(),
			"Result",
			fmt.Sprintf("cdp_%s.txt", entry.Costco.CatId),
		)
		c.scheduler.Repeat(ctx, interval, "Fetch "+entry.Cat.Name, func() {
			itemModel, err := costco.FetchItemModel(
				c.scope,
				entry.Cat.Name,
				entry.Costco.ItemId,
				entry.Costco.CategoryId,
				entry.Costco.ProductId,
				entry.Costco.QueryStringPatch,
			)
			if ezgo.IsErr(err) {
				ezgo.LogCausesf(c.scope.GetLogger(), err, "FetchItemModel(%s)", entry.Cat.Name)
				return
			}

			resultLogger.Info(
				"Fetched cdp",
				zap.String("name", entry.Cat.Name),
				zap.Float64("price", itemModel.Price),
				zap.Bool("inStock", itemModel.Available),
			)
		})
	}
}

func (c *Crawler) Terminate() {
	c.scheduler.Terminate()
}
