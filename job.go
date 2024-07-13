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

func crawl(feedID int, url string, needNotify bool) {
	rawFeed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return
	}
	slog.Debug("Fetch feed", "url", url)

	for _, item := range rawFeed.Items {
		if item == nil {
			break
		}

		title := item.Title
		link := item.Link
		publishedAt := *item.PublishedParsed

		exists, err := models.Articles(
			models.ArticleWhere.FeedID.EQ(feedID),
			models.ArticleWhere.PublishedAt.EQ(publishedAt),
		).ExistsG()
		if err != nil {
			slog.Error("Error occurred", "error", err)
			continue
		}
		if exists {
			slog.Debug("Article already exists", "feedID", feedID, "title", title)
			continue
		}

		article := &models.Article{
			FeedID:      feedID,
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
		slog.Debug("Create article", "feedID", feedID, "title", title)
	}
}

func crawlingJob() {
	slog.Debug("Start Crawling")

	feeds, err := models.Feeds().AllG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return
	}

	for _, feed := range feeds {
		crawl(feed.FeedID, feed.URL, true)
		time.Sleep(1 * time.Second)
	}

	slog.Debug("Finish Crawling")
}

func notify(channelID string, msg string) error {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		return err
	}
	return nil
}

func notificationJob() {
	slog.Debug("Start Notification")

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

	slog.Debug("Finish Notification")
}
