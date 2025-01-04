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
	crawler.Kickoff(ctx)

	time.Sleep(24 * time.Hour)

	crawler.Terminate()
}
