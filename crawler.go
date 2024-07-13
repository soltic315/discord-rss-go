package main

import (
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
)

func fetchRawFeed(url string) (*gofeed.Feed, error) {
	rawFeed, err := gofeed.NewParser().ParseURL(url)
	return rawFeed, err
}

func crawl() {
	slog.Debug("Start Crawling")

	feeds, err := getFeeds()
	if err != nil {
		slog.Error("error", err)
		return
	}

	for _, feed := range feeds {
		feedID := feed.FeedID
		url := feed.URL

		rawFeed, err := fetchRawFeed(url)
		if err != nil {
			slog.Error("error", err)
			continue
		}
		slog.Debug("Fetch feed", "url", url)

		for _, item := range rawFeed.Items {
			if item == nil {
				break
			}

			title := item.Title
			link := item.Link
			publishedAt := *item.PublishedParsed

			ok, err := existsArticle(feedID, publishedAt)
			if err != nil {
				slog.Error("error", err)
				continue
			}
			if ok {
				slog.Debug("Article already exists", "feedID", feedID, "title", title)
				continue
			}

			msg := fmt.Sprintf("**📰 | %s**\n%s", title, link)
			err = sendMessage(ChannelID, msg)
			if err != nil {
				slog.Error("error", err)
			}
			slog.Info("Notify article", "feedID", feedID, "title", title)

			_, err = createArticle(feedID, title, link, publishedAt)

			if err != nil {
				slog.Error("error", err)
				continue
			}
			slog.Debug("Create article", "feedID", feedID, "title", title)
		}
	}

	slog.Debug("Finish Crawling")
}
