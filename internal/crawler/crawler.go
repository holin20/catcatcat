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

func (c *Crawler) Kickoff(ctx context.Context) {
	items := []struct {
		name         string
		priceUrl     string
		inventoryUrl string
		interval     time.Duration

		productId        string
		itemId           string
		categoryId       string
		queryStringPatch string
	}{
		{
			name:             "Macbook Air M3 15\" 16G 256GB",
			priceUrl:         "https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620577642&catalogId=10701&productId=3074457345620577640&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
			productId:        "3074457345620577640",
			itemId:           "3074457345620577642",
			categoryId:       "10701",
			queryStringPatch: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
		{
			name:             "CurrentBody Skin LED Mask Face & Neck Kit",
			inventoryUrl:     "https://www.costco.com/AjaxGetInventoryDetail?itemId=3074457345620439817&catalogId=10701&productId=3074457345620439815&WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false&zipCode=98034",
			productId:        "3074457345620439815",
			itemId:           "3074457345620439817",
			categoryId:       "10701",
			queryStringPatch: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
	}

	for _, item := range items {
		item := item
		interval := item.interval
		if interval == 0 {
			interval = 5 * time.Minute
		}

		resultLogger := c.scope.GetLogger().Named("Result")
		c.scheduler.Repeat(ctx, interval, "Fetch "+item.name, func() {
			itemModel, err := costco.FetchItemModel(c.scope, item.itemId, item.categoryId, item.productId, item.queryStringPatch)
			if ezgo.IsErr(err) {
				ezgo.LogCausesf(c.scope.GetLogger(), err, "FetchItemModel(%s)", item.name)
				return
			}

			resultLogger.Info(
				"Fetched Item Info",
				zap.String("item", item.name),
				zap.Float64("price", itemModel.Price),
				zap.Bool("available", itemModel.Available),
			)
		})
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

func (c *Crawler) Terminate() {
	c.scheduler.Terminate()
}
