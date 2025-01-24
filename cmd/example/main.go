package main

import (
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
	orm "github.com/holin20/catcatcat/pkg/ezgo/orm"
)

func main() {
	actualizeTest()

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

func actualizeTest() {
	type Dog struct {
		Name    string `sql:"name"`
		Species int    `sql:"species"`
		IsAlive bool   `sql:"is_alive"`
	}

	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	structTagSql := ezgo.NewStructTag[Dog]("sql")

	//orm.Actualize(db, structTagSql)

	// dog1 := Dog{
	// 	Name:    "lalala",
	// 	Species: 45,
	// 	IsAlive: true,
	// }
	// err = orm.Create(db, &dog1, structTagSql)
	// ezgo.AssertNoError(err, "ezgo.Create")

	dogs, err := orm.Load(db, structTagSql, 6)
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
