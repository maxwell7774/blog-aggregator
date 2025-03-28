package main

import (
	"context"
	"fmt"

	"github.com/maxwell7774/blog-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
    /*
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %s <feedURL>", cmd.Name)
    }

    feedURL := cmd.Args[0]
    */
    feedURL := "https://www.wagslane.dev/index.xml"

    feed, err := rss.FetchFeed(context.Background(), feedURL)
    if err != nil {
        return fmt.Errorf("error fetching feed: %w", err)
    }

    fmt.Printf("Successfully retrieved feed: %v\n", feed)

    return nil
}
