package main

import (
	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/holin20/catcatcat/pkg/ezgo/orm"
)

var (
	CATS = []*schema.Cat{
		{
			CatId: "0",
			Name:  "Macbook Air M3 15\" 16G 256GB",
		},
		{
			CatId: "1",
			Name:  "CurrentBody Skin LED Mask Face And Neck Kit",
		},
	}

	COSTCO_FETCHER_LIST = []*schema.CostcoFetcher{
		{
			CatId:      "0",
			ItemId:     "3074457345620577642",
			ProductId:  "3074457345620577640",
			CategoryId: "10701",

			QueryStringPath: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
		{
			CatId:      "1",
			ItemId:     "3074457345620439817",
			ProductId:  "3074457345620439815",
			CategoryId: "10701",
			Name:       "CurrentBody Skin LED Mask Face & Neck Kit",

			QueryStringPath: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
	}
)

func main() {
	actualData()
}

func actualizeSchemas() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")

	err = orm.Actualize(db, orm.NewSchema[schema.Cat]())
	ezgo.AssertNoError(err, "actualize cat")

	err = orm.Actualize(db, orm.NewSchema[schema.Cdp]())
	ezgo.AssertNoError(err, "actualize cdp")

	err = orm.Actualize(db, orm.NewSchema[schema.CostcoFetcher]())
	ezgo.AssertNoError(err, "actualize CostcoFetcher")
}

func actualData() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")

	for _, cat := range CATS {
		err := orm.Create(db, orm.NewSchema[schema.Cat](), cat)
		ezgo.AssertNoError(err, "creating cat")
	}

	for _, fetcher := range COSTCO_FETCHER_LIST {
		err := orm.Create(db, orm.NewSchema[schema.CostcoFetcher](), fetcher)
		ezgo.AssertNoError(err, "creating cat")
	}
}
