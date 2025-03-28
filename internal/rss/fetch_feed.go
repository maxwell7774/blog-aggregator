package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "gator")

    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    dat, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var feed RSSFeed
    err = xml.Unmarshal(dat, &feed)
    if err != nil {
        return nil, err
    }

    feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
    feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

    for i, item := range feed.Channel.Item {
        feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
        feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
    }

    return &feed, nil
}
