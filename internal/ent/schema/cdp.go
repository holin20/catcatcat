package schema

import "github.com/holin20/catcatcat/pkg/ezgo/orm"

var CdpSchema = orm.NewSchema[Cdp]()

type Cdp struct {
	CatId   string  `sql:"cat_id" pk:"true"`
	Price   float64 `sql:"price"`
	InStock bool    `sql:"in_stock"`
}
