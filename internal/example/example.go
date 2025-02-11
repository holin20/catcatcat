package example

import "github.com/holin20/catcatcat/internal/model"

var (
	WATCH_LIST = []*model.Watch{
		{
			WatchId:  "0",
			CatId:    "0",
			NickName: "Macbook Air M3 15\" 16G 256GB",
		},
		{
			WatchId:  "1",
			CatId:    "1",
			NickName: "CurrentBody Skin LED Mask Face And Neck Kit",
		},
		{
			WatchId:  "2",
			CatId:    "2",
			NickName: `MacBook Pro Laptop (14-inch) - Apple M3 Pro Chip, 11-core CPU, 14-core GPU, 18GB Memory, 512GB SSD Storage`,
		},
	}

	CATS = []*model.Cat{
		{
			CatId: "0",
			Name:  "Macbook Air M3 15\" 16G 256GB",
		},
		{
			CatId: "1",
			Name:  "CurrentBody Skin LED Mask Face And Neck Kit",
		},
		{
			CatId: "2",
			Name:  `MacBook Pro Laptop (14-inch) - Apple M3 Pro Chip, 11-core CPU, 14-core GPU, 18GB Memory, 512GB SSD Storage`,
		},
	}

	COSTCO_FETCHER_LIST = []*model.CostcoFetcher{
		{
			CatId:      "0",
			ItemId:     "3074457345620577642",
			ProductId:  "3074457345620577640",
			CategoryId: "10701",

			QueryStringPatch: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
		{
			CatId:      "1",
			ItemId:     "3074457345620439817",
			ProductId:  "3074457345620439815",
			CategoryId: "10701",
			Name:       "CurrentBody Skin LED Mask Face & Neck Kit",

			QueryStringPatch: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
		{
			CatId:      "2",
			ItemId:     "3074457345620221936",
			ProductId:  "3074457345620223430",
			CategoryId: "10701",
			Name:       `MacBook Pro Laptop (14-inch) - Apple M3 Pro Chip, 11-core CPU, 14-core GPU, 18GB Memory, 512GB SSD Storage`,

			QueryStringPatch: "WH=1250-3pl%2C1321-wm%2C1456-3pl%2C283-wm%2C561-wm%2C725-wm%2C731-wm%2C758-wm%2C759-wm%2C847_0-cor%2C847_0-cwt%2C847_0-edi%2C847_0-ehs%2C847_0-membership%2C847_0-mpt%2C847_0-spc%2C847_0-wm%2C847_1-cwt%2C847_1-edi%2C847_d-fis%2C847_lg_n1f-edi%2C847_NA-cor%2C847_NA-pharmacy%2C847_NA-wm%2C847_ss_u362-edi%2C847_wp_r458-edi%2C951-wm%2C952-wm%2C9847-wcs%2C115-bd&warehouse=1-wh&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false",
		},
	}

	CRAWL_LIST = []struct {
		CatId  string
		Cat    *model.Cat
		Costco *model.CostcoFetcher
	}{
		{
			CatId:  "0",
			Cat:    CATS[0],
			Costco: COSTCO_FETCHER_LIST[0],
		},
		{
			CatId:  "1",
			Cat:    CATS[1],
			Costco: COSTCO_FETCHER_LIST[1],
		},
		{
			CatId:  "2",
			Cat:    CATS[2],
			Costco: COSTCO_FETCHER_LIST[2],
		},
	}
)
