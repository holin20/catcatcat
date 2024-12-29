package main

import (
	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"

	"go.uber.org/zap"
)

func main() {
	scope := ezgo.Must(ezgo.NewScopeWithDefaultLogger())
	defer scope.Close()

	run(scope)
}

func run(scope *ezgo.Scope) {
	price, err := costco.FetchMacbookPrice()
	if ezgo.IsErr(err) {
		ezgo.LogCauses(scope.GetLogger(), err, "FetchMacbookPrice")
		return
	}

	scope.GetLogger().Info("Price", zap.Float64("price", price))
}
