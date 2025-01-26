package schema

type Cdp struct {
	CatId   string  `sql:"cat_id"`
	Price   float64 `sql:"price"`
	InStock bool    `sql:"in_stock"`
}
