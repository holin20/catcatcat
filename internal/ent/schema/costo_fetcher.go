package schema

import "github.com/holin20/catcatcat/pkg/ezgo/orm"

var CostcoFetcherSchema = orm.NewSchema[CostcoFetcher]()

type CostcoFetcher struct {
	CatId string `sql:"cat_id" unique:"true"`

	ProductId  string `sql:"product_id"`
	ItemId     string `sql:"item_id"`
	CategoryId string `sql:"category_id"`

	QueryStringPatch string `sql:"query_string_patch"`
}
