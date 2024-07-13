package main

import (
	"discord-rss-go/models"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func getOrCreateFeed(url string, title string) (*models.Feed, error) {
	var feed *models.Feed

	exists, err := models.Feeds(
		models.FeedWhere.URL.EQ(url),
	).ExistsG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return nil, err
	}
	if exists {
		feed, err = models.Feeds(
			models.FeedWhere.URL.EQ(url),
		).OneG()
		if err != nil {
			slog.Error("Error occurred", "error", err)
			return nil, err
		}
	} else {
		feed = &models.Feed{
			URL:   url,
			Title: title,
		}
		err = feed.InsertG(boil.Infer())
		if err != nil {
			slog.Error("Error occurred", "error", err)
			return nil, err
		}
		slog.Info("Create feed", "url", feed.URL)
	}

	return feed, nil
}

func feedSubscribeCommand(url string, channelID string) string {
	var feed *models.Feed

	exists, err := models.Subscriptions(
		qm.InnerJoin("feeds ON subscriptions.feed_id = feeds.feed_id"),
		qm.Where("subscriptions.channel_id = ?", channelID),
		qm.Where("feeds.url = ?", url),
	).ExistsG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	if exists {
		return fmt.Sprintf("このチャンネルにはすでにそのフィード(%s)が登録されています", url)
	}

	rawFeed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	slog.Debug("Fetch feed", "url", url)

	feed, err = getOrCreateFeed(url, rawFeed.Title)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}

	subscription := &models.Subscription{
		FeedID:    feed.FeedID,
		ChannelID: channelID,
	}
	err = subscription.InsertG(boil.Infer())
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	slog.Info("Create subscription", "feedID", feed.FeedID, "ChannelID", channelID)

	content := fmt.Sprintf("フィード[%s](%s)の購読が完了しました\n", rawFeed.Title, url)
	if len(rawFeed.Items) != 0 {
		content += fmt.Sprintf("> **📰 | %s**\n%s", rawFeed.Items[0].Title, rawFeed.Items[0].Link)
	}

	return content
}

func feedListCommand(channelID string) string {
	var status string

	subscriptions, err := models.Subscriptions(
		models.SubscriptionWhere.ChannelID.EQ(channelID),
		qm.Load(models.SubscriptionRels.Feed),
	).AllG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	if subscriptions == nil {
		return "フィードがありません"
	}

	content := ""
	for _, subscription := range subscriptions {
		feed := subscription.R.Feed

		if feed.RequestFailureCount <= CrawlingStopFailureCount {
			status = "✅"
		} else {
			status = "🚫"
		}

		content += fmt.Sprintf("%s %s (ID: %d, URL: %s)\n", status, feed.Title, subscription.SubscriptionID, feed.URL)
	}

	return content
}

func feedRemoveCommand(subscriptionID int) string {
	subscription, err := models.FindSubscriptionG(subscriptionID)
	if subscription == nil {
		return fmt.Sprintf("フィード(ID: %d)を見つけられませんでした", subscriptionID)
	}
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}

	_, err = subscription.DeleteG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	slog.Info("Delete feed", "subscriptionID", subscriptionID)

	feed, err := models.FindFeedG(subscription.FeedID)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}

	// TODO: Delete feed if no subscription

	return fmt.Sprintf("フィード[%s](%s) (ID: %d)が削除されました", feed.Title, feed.URL, subscription.SubscriptionID)
}
