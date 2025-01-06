package model

type CatDataPoint struct {
	CatId   string
	Price   float64
	InStock bool
}

type Cat struct {
	CatId string
	Name  string
}
