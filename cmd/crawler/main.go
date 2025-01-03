package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/crawler"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger()
	defer scope.Close()

	ctx := context.Background()

	crawler := crawler.NewCrawler(scope)
	crawler.Start(ctx)

	time.Sleep(24 * time.Hour)

	ezgo.AsyncVoid(crawler.Terminate).Await()
}
