package crawler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/holin20/catcatcat/pkg/ezgo/orm"
	"go.uber.org/zap"
)

const (
	defaultCrawlInterval = 1 * time.Hour
)

var (
	cdpSchema           = orm.NewSchema[schema.Cdp]()
	catSchema           = orm.NewSchema[schema.Cat]()
	costcoFetcherSchema = orm.NewSchema[schema.CostcoFetcher]()
)

type CrawlListEntry = struct {
	CatId  string
	Cat    *schema.Cat
	Costco *schema.CostcoFetcher
}

type Crawler struct {
	scheduler    *ezgo.Scheduler
	scope        *ezgo.Scope
	crawlList    []CrawlListEntry
	crawInterval time.Duration
	db           *ezgo.PostgresDB
}

func NewCrawler(scope *ezgo.Scope) *Crawler {
	scope = scope.WithLogger(scope.GetLogger().Named("Crawler"))
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "failed to open pdb")
	return &Crawler{
		scheduler:    ezgo.NewScheduler(scope),
		crawInterval: defaultCrawlInterval,
		scope:        scope,
		db:           db,
	}
}

func (c *Crawler) WithCrawlListFromDB() *Crawler {
	cats, err := orm.Load(c.db, catSchema)
	ezgo.AssertNoErrorf(err, "load cat schema")
	c.scope.GetLogger().Info("Cats created", zap.Int("count", len(cats)))

	costcoFetchers, err := orm.Load(c.db, costcoFetcherSchema)
	ezgo.AssertNoErrorf(err, "load costco fetcher schema")
	c.scope.GetLogger().Info("Costco fetchers created", zap.Int("count", len(costcoFetchers)))

	for _, cat := range cats {
		// Find costco fetcher for this cat
		var f *schema.CostcoFetcher
		for _, cf := range costcoFetchers {
			if cat.CatId == cf.CatId {
				f = cf
				break
			}
		}
		if f == nil {
			c.scope.GetLogger().Error(
				"Cat id has no fetcher",
				zap.String("cat_id", cat.CatId),
				zap.String("cat_name", cat.Name),
			)
			continue
		}
		if f.Disabled {
			c.scope.GetLogger().Info(
				"This fetcher is disabled",
				zap.String("cat_id", cat.CatId),
				zap.String("cat_name", cat.Name),
			)
			continue
		}

		entry := CrawlListEntry{
			CatId:  cat.CatId,
			Cat:    cat,
			Costco: f,
		}
		c.crawlList = append(c.crawlList, entry)

		c.scope.GetLogger().Info(
			"Crawl list is created",
			zap.String("cat_id", entry.CatId),
			zap.String("cat_name", entry.Cat.Name),
		)
	}

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
			// parallel version has potential concurrent map read/write because
			// the 2 json fetch share the same ezgo http wrapper.
			itemModel, err := costco.FetchItemModelSequential(
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
				zap.String("cat_id", entry.Cat.CatId),
				zap.Float64("price", itemModel.Price),
				zap.Bool("inStock", itemModel.Available),
			)

			c.writeCdp(entry.Cat.CatId, itemModel.Price, itemModel.Available)
		}
	})
}

func (c *Crawler) writeCdp(catId string, price float64, inStock bool) {
	err := orm.Create[schema.Cdp](c.db, cdpSchema, &schema.Cdp{
		CatId:   catId,
		Price:   price,
		InStock: inStock,
	})

	if ezgo.IsErr(err) {
		ezgo.LogCauses(c.scope.GetLogger(), err, "failed to write cdp")
	}
}

func (c *Crawler) Terminate() {
	c.scheduler.Terminate()
}
