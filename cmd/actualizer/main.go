package main

import (
	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/holin20/catcatcat/pkg/ezgo/orm"
)

func main() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")

	// err = orm.Actualize(db, orm.NewSchema[schema.Cat]())
	// ezgo.AssertNoError(err, "actualize cat")

	// err = orm.Actualize(db, orm.NewSchema[schema.Cdp]())
	// ezgo.AssertNoError(err, "actualize cdp")

	// err = orm.Actualize(db, orm.NewSchema[schema.CostcoFetcher]())
	// ezgo.AssertNoError(err, "actualize CostcoFetcher")

	err = orm.Create(db, orm.NewSchema[schema.Cat](), &schema.Cat{
		CatId: "0",
		Name:  "MAcbook",
	})

	ezgo.AssertNoError(err, "orm.Create cat")
}
