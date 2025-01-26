package schema

type CostcoFetcher struct {
	CatId string `sql:"cat_id" unique:"true"`

	ProductId  string `sql:"product_id"`
	ItemId     string `sql:"item_id"`
	CategoryId string `sql:"category_id"`

	Name string `sql:"name"`

	QueryStringPath string `sql:"query_string_path"`
}
