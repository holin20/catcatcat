package main

import (
	"fmt"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	err := ezgo.GmailSender().
		From("catcatcattm@gmail.com").
		To("holin20@gmail.com").
		Subject("Subject").
		Body("Body").
		PasswordEnv("CATCATCAT_GAPP_PW").
		Send()

	if ezgo.IsErr(err) {
		fmt.Println(err)
	}
}
