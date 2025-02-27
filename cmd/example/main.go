package main

import (
	"context"
	"fmt"
	"time"

	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/internal/monitor"
	"github.com/holin20/catcatcat/pkg/ezgo"
	orm "github.com/holin20/catcatcat/pkg/ezgo/orm"
)

func main() {
	//ormCreateTest()
	//ormLoadTest()

	//testLoadLastN()

	//testEntCdpQuery()

	//	PostgresSqlQueryTest()
	// err := ezgo.GmailSender().
	// 	From("catcatcattm@gmail.com").
	// 	To("holin20@gmail.com").
	// 	Subject("Subject").
	// 	Body("Body").
	// 	PasswordEnv("CATCATCAT_GAPP_PW").
	// 	Send()

	// if ezgo.IsErr(err) {
	// 	fmt.Println(err)
	// }

	// db, err := ezgo.NewLocalPostgresDB(
	// 	"postgres",
	// 	"postgres",
	// 	54320,
	// 	"postgres",
	// )

	// if ezgo.IsErr(err) {
	// 	ezgo.AssertNoError(err, "db open")
	// }

	// err = db.Insert("cat_test", map[string]*ezgo.SqlCol{
	// 	"name": ezgo.SqlColString("mac"),
	// })
	// if ezgo.IsErr(err) {
	// 	ezgo.AssertNoError(err, "db open")
	// }

	// sql := ezgo.Must(ezgo.NewSqlBuilder().Select("id, name").From("cat_test").Build())
	// colNames, resultSets, err := db.Query(sql)
	// ezgo.AssertNoError(err, "db query: "+sql)
	// fmt.Println(colNames)
	// fmt.Println(resultSets)
}

func testTableExists() {
	// db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	// ezgo.AssertNoError(err, "NewLocalPostgresDB")
	// defer db.Close()

	// exists, err := orm.TableExists(db, "cdp")
	// ezgo.AssertNoError(err, "TableExists")
	// fmt.Printf("exist: %d\n", ezgo.If(exists, 1, 0))
}

// func testAlterSchema() {
// 	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
// 	ezgo.AssertNoError(err, "NewLocalPostgresDB")
// 	defer db.Close()

// 	err = orm.alterSchema(db, schema.CostcoFetcherSchema)
// 	ezgo.AssertNoError(err, "GetTableColumns")
// }

// func testGetTableColumns() {
// 	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
// 	ezgo.AssertNoError(err, "NewLocalPostgresDB")
// 	defer db.Close()

// 	colNameToType, err := orm.getTableColumns(db, "cdp")
// 	ezgo.AssertNoError(err, "GetTableColumns")
// 	fmt.Println(ezgo.ToJsonString(colNameToType))
// }

func testEntCdpQuery() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	cdp, _, err := monitor.NewEntCdpQuery(db, "0", monitor.CdpInStock).Query(context.Background(), time.Now())
	ezgo.AssertNoError(err, "NewEntCdpQuery")
	fmt.Println(ezgo.ToJsonString(cdp))
}

func testLoadLastN() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	results, err := orm.LoadLastN(db, orm.NewSchema[schema.Cdp](), &schema.Cdp{CatId: "1"}, 1)
	ezgo.AssertNoError(err, "LoadLastN")

	for _, v := range results {
		_, y := v.Unpack()
		fmt.Println(ezgo.ToJsonString(y))
	}
}

type Dog struct {
	Name    string `sql:"name"`
	Species int    `sql:"species"`
	IsAlive bool   `sql:"is_alive"`
}

func ormActualizeTest() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	dogSchema := orm.NewSchema[Dog]()

	orm.Actualize(db, dogSchema)
}

func ormCreateTest() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	dogSchema := orm.NewSchema[Dog]()
	orm.Actualize(db, dogSchema)

	dog1 := Dog{
		Name:    "lalala",
		Species: 45,
		IsAlive: true,
	}
	err = orm.Create(db, dogSchema, &dog1)
	ezgo.AssertNoError(err, "ezgo.Create")
}

func ormLoadTest() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	dogSchema := orm.NewSchema[Dog]()

	//orm.Actualize(db, structTagSql)

	// dog1 := Dog{
	// 	Name:    "lalala",
	// 	Species: 45,
	// 	IsAlive: true,
	// }
	// err = orm.Create(db, &dog1, structTagSql)
	// ezgo.AssertNoError(err, "ezgo.Create")

	dogs, err := orm.Load(db, dogSchema, 2)
	ezgo.AssertNoError(err, "orm.LoadFrom")

	fmt.Println(ezgo.ToJsonString(dogs))
}

func PostgresSqlQueryTest() {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()
	// q := monitor.NewPostgresSqlQuery(
	// 	db,
	// 	"select ts, price from cdp order by ts desc limit 1",
	// 	"ts",
	// 	"price",
	// )

	// q.Query(context.Background(), time.Now())

	_ = db.Insert("cdp", map[string]*ezgo.SqlCol{
		"cat_id":   ezgo.SqlColString("0"),
		"ts":       ezgo.SqlColInt(time.Now().Unix()),
		"price":    ezgo.SqlColFloat(1099.0),
		"in_stock": ezgo.SqlColBool(true),
	})

	_ = db.Insert("cat", map[string]*ezgo.SqlCol{
		"cat_id": ezgo.SqlColString("0"),
		"name":   ezgo.SqlColInt(time.Now().Unix()),
		"ts":     ezgo.SqlColInt(time.Now().Unix()),
	})

	ezgo.AssertNoError(err, "db.Insert")
}
