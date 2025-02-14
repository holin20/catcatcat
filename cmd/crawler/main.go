package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/crawler"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger("Crawler")
	defer scope.Close()

	ctx := context.Background()

	crawler := crawler.NewCrawler(scope).
		WithCrawlListFromDB().
		WithCrawlInterval(1 * time.Hour)

	crawler.Kickoff(ctx)

	time.Sleep(24 * time.Hour * 7)

	crawler.Terminate()
}
