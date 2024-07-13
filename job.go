package main

import (
	"discord-rss-go/models"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func crawl(feed *models.Feed, needNotify bool) {
	rawFeed, err := gofeed.NewParser().ParseURL(feed.URL)
	if err == nil {
		feed.RequestFailureCount = 0
		_, updateErr := feed.UpdateG(boil.Infer())
		if updateErr != nil {
			slog.Error("Error occurred", "error", updateErr)
			return
		}

		slog.Debug("Fetch feed", "url", feed.URL)
	} else {
		feed.RequestFailureCount = feed.RequestFailureCount + 1
		_, updateErr := feed.UpdateG(boil.Infer())
		if updateErr != nil {
			slog.Error("Error occurred", "error", updateErr)
			return
		}

		slog.Error("Failed to fetch feed", "error", err, "requestFailureCount", feed.RequestFailureCount)
		return
	}

	for _, item := range rawFeed.Items {
		title := item.Title
		link := item.Link
		publishedAt := *item.PublishedParsed

		exists, err := models.Articles(
			models.ArticleWhere.FeedID.EQ(feed.FeedID),
			models.ArticleWhere.PublishedAt.EQ(publishedAt),
		).ExistsG()
		if err != nil {
			slog.Error("Error occurred", "error", err)
			continue
		}
		if exists {
			slog.Debug("Article already exists", "feedID", feed.FeedID, "title", title)
			continue
		}

		article := &models.Article{
			FeedID:      feed.FeedID,
			Title:       title,
			Link:        link,
			NeedNotify:  needNotify,
			PublishedAt: publishedAt,
		}
		err = article.InsertG(boil.Infer())

		if err != nil {
			slog.Error("Error occurred", "error", err)
			continue
		}
		slog.Info("Create article", "articleId", article.ArticleID)
	}
}

func crawlingJob() {
	slog.Info("Start Crawling")

	feeds, err := models.Feeds(
		qm.InnerJoin("subscriptions ON feeds.feed_id = subscriptions.feed_id"),
		qm.Where("feeds.request_failure_count <= ?", CrawlingStopFailureCount),
	).AllG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return
	}

	for _, feed := range feeds {
		needNotify, err := models.Articles(
			qm.Where("feed_id = ?", feed.FeedID),
		).ExistsG()
		if err != nil {
			slog.Error("Error occurred", "error", err)
			return
		}

		crawl(feed, needNotify)
		time.Sleep(1 * time.Second)
	}

	slog.Info("Finish Crawling")
}

func notify(channelID string, msg string) error {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		return err
	}
	return nil
}

func notificationJob() {
	slog.Info("Start Notification")

	articles, err := models.Articles(
		models.ArticleWhere.NeedNotify.EQ(true),
	).AllG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return
	}

	for _, article := range articles {
		msg := fmt.Sprintf("**📰 | %s**\n%s", article.Title, article.Link)

		subscriptions, err := models.Subscriptions(
			qm.InnerJoin("feeds ON subscriptions.feed_id = feeds.feed_id"),
			qm.Where("feeds.feed_id = ?", article.FeedID),
		).AllG()
		if err != nil {
			slog.Error("Error occurred", "error", err)
			return
		}
		for _, subscription := range subscriptions {
			err = notify(subscription.ChannelID, msg)
			time.Sleep(1 * time.Second)
			if err != nil {
				slog.Error("Error occurred", "error", err)
			}
			slog.Info("Notify article", "title", article.Title, "channelID", subscription.ChannelID)
		}

		article.NeedNotify = false
		_, err = article.UpdateG(boil.Infer())
		if err != nil {
			slog.Error("Error occurred", "error", err)
		}
	}

	slog.Info("Finish Notification")
}
