package main

import (
	"fmt"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	price, err := costco.FetchMacbookPrice()
	if ezgo.IsErr(err) {
		panic(err)
	}
	fmt.Println(price)
}
