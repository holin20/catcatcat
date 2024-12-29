package main

import (
	"fmt"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/gen"
)

func main() {
	price, err := costco.FetchMacbookPrice()
	if gen.IsErr(err) {
		panic(err)
	}
	fmt.Println(price)
}
