package main

import (
	"fmt"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	price, err := costco.FetchMacbookPrice()
	if ezgo.IsErr(err) {
		fmt.Println(ezgo.NewCause(err, "FetchMacbookPrice"))
		return
	}
	fmt.Println(price)
}
