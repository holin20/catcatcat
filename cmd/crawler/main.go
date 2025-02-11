package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/crawler"
	"github.com/holin20/catcatcat/internal/example"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger("Crawler")
	defer scope.Close()

	ctx := context.Background()

	crawler := crawler.NewCrawler(scope).
		WithCrawlList(example.CRAWL_LIST).
		WithCrawlInterval(1 * time.Hour)

	crawler.Kickoff(ctx)

	time.Sleep(27 * time.Hour * 7)

	crawler.Terminate()
}
