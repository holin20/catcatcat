package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/fetcher/costco"
	"github.com/holin20/catcatcat/pkg/ezgo"

	"go.uber.org/zap"
)

func main() {
	scope := ezgo.Must(ezgo.NewScopeWithDefaultLogger())
	defer scope.Close()

	scheduler := ezgo.NewScheduler(scope)

	ctx := context.Background()
	scheduler.RepeatN(ctx, 10*time.Second, 6, ezgo.NewNamedTask("FetchMacbookPrice", func() {
		price, err := costco.FetchMacbookPrice()
		if ezgo.IsErr(err) {
			ezgo.LogCauses(scope.GetLogger(), err, "FetchMacbookPrice")
			return
		}
		scope.GetLogger().Info("Price", zap.Float64("price", price))
	}))

	scheduler.Join()
}

// func run(scope *ezgo.Scope) {
// 	price, err := costco.FetchMacbookPrice()
// 	if ezgo.IsErr(err) {
// 		ezgo.LogCauses(scope.GetLogger(), err, "FetchMacbookPrice")
// 		return
// 	}

// 	scope.GetLogger().Info("Price", zap.Float64("price", price))
// }
