package schema

import "github.com/holin20/catcatcat/pkg/ezgo/orm"

var CatSchema = orm.NewSchema[Cat]()

type Cat struct {
	CatId string `sql:"cat_id" unique:"true"`
	Name  string `sql:"name"`
}
