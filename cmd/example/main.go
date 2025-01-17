package main

import (
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
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

	db, err := ezgo.NewLocalPostgresDB(
		"postgres",
		"postgres",
		54320,
		"postgres",
	)

	if ezgo.IsErr(err) {
		ezgo.AssertNoError(err, "db open")
	}

	// err = db.Insert("cat_test", map[string]*ezgo.SqlCol{
	// 	"name": ezgo.SqlColString("mac"),
	// })
	// if ezgo.IsErr(err) {
	// 	ezgo.AssertNoError(err, "db open")
	// }

	sql := ezgo.Must(ezgo.NewSqlBuilder().Select("id, name").From("cat_test").Build())
	colNames, resultSets, err := db.Query(sql)
	ezgo.AssertNoError(err, "db query: "+sql)
	fmt.Println(colNames)
	fmt.Println(resultSets)
}
