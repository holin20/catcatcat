package schema

type Cat struct {
	CatId string `sql:"cat_id" unique:"true"`
	Name  string `sql:"name"`
}
