package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/crawler"
	"github.com/holin20/catcatcat/internal/example"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger()
	defer scope.Close()

	ctx := context.Background()

	crawler := crawler.NewCrawler(scope).
		WithCrawlList(example.CRAWL_LIST).
		WithCrawlInterval(1 * time.Minute)

	crawler.Kickoff(ctx)

	time.Sleep(24 * time.Hour)

	crawler.Terminate()
}
