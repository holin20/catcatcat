package crawler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/internal/model"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

const (
	defaultCrawlInterval = 1 * time.Hour
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
		scheduler:    ezgo.NewScheduler(scope),
		crawInterval: defaultCrawlInterval,
		scope:        scope,
	}
}

func (c *Crawler) WithCrawlList(crawlList []CrawlListEntry) *Crawler {
	c.crawlList = crawlList
	return c
}

func (c *Crawler) WithCrawlInterval(crawlInterval time.Duration) *Crawler {
	if crawlInterval > 0 {
		c.crawInterval = crawlInterval
	}
	return c
}

func (c *Crawler) Kickoff(ctx context.Context) {
	c.scope.GetLogger().Info("Kicking off crawler!", zap.Duration("crawl_interval", c.crawInterval))

	resultLoggers := ezgo.SliceApply(c.crawlList, func(i int, entry CrawlListEntry) *zap.Logger {
		return ezgo.CloneLogger(
			c.scope.GetLogger(),
			"Result",
			fmt.Sprintf("cdp_%s.txt", entry.Costco.CatId),
		)
	})

	goHttpClient := &http.Client{
		Jar: ezgo.Arg1(cookiejar.New(nil)),
	}
	c.scheduler.Repeat(ctx, c.crawInterval, "Crawler Single Fetcher", func() {
		for i, entry := range c.crawlList {
			itemModel, err := costco.FetchItemModel(
				c.scope,
				ezgo.NewHttpClientWithCustomClient(goHttpClient, true),
				entry.Cat.Name,
				entry.Costco.ItemId,
				entry.Costco.CategoryId,
				entry.Costco.ProductId,
				entry.Costco.QueryStringPatch,
			)
			if ezgo.IsErr(err) {
				ezgo.LogCausesf(c.scope.GetLogger(), err, "FetchItemModel(%s)", entry.Cat.Name)
				continue
			}

			resultLoggers[i].Info(
				"Fetched cdp",
				zap.String("name", entry.Cat.Name),
				zap.Float64("price", itemModel.Price),
				zap.Bool("inStock", itemModel.Available),
			)
		}
	})
}

func (c *Crawler) Terminate() {
	c.scheduler.Terminate()
}
