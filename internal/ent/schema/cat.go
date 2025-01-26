package schema

type Cat struct {
	CatId string `sql:"cat_id"`
	Name  string `sql:"name"`
}
