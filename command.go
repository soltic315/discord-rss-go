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

func createFeedAndCrawl(url string) (*models.Feed, error) {
	rawFeed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return nil, err
	}
	slog.Debug("Fetch feed", "url", url)

	feed := &models.Feed{
		URL:   url,
		Title: rawFeed.Title,
	}
	err = feed.InsertG(boil.Infer())
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return nil, err
	}
	slog.Info("Create feed", "url", feed.URL)

	crawl(feed.FeedID, feed.URL, false)

	return feed, nil
}

func feedSubscribeCommand(url string, channelID string, userName string) string {
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
		return fmt.Sprintf("このチャンネルにはすでにそのフィード（%s）が登録されています", url)
	}

	exists, err = models.Feeds(
		models.FeedWhere.URL.EQ(url),
	).ExistsG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	if !exists {
		_, err = createFeedAndCrawl(url)
		if err != nil {
			slog.Error("Error occurred", "error", err)
			return "フィードの取得中に不具合が発生しました"
		}
	}

	feed, err = models.Feeds(
		models.FeedWhere.URL.EQ(url),
	).OneG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}

	subscription := &models.Subscription{
		FeedID:    feed.FeedID,
		ChannelID: channelID,
		CreatedBy: userName,
	}
	err = subscription.InsertG(boil.Infer())
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}
	slog.Info("Create subscription", "feedID", feed.FeedID, "ChannelID", channelID)

	article, err := models.Articles(
		qm.InnerJoin("feeds ON articles.feed_id = feeds.feed_id"),
		qm.Where("feeds.url = ?", url),
	).OneG()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		return "フィードの取得中に不具合が発生しました"
	}

	return fmt.Sprintf("購読フィード: %s\n**📰 | %s**\n%s", feed.Title, article.Title, article.Link)
}

func feedListCommand(channelID string) string {
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

		content += fmt.Sprintf("ID: %d - タイトル: %s\n", subscription.SubscriptionID, feed.Title)
		content += fmt.Sprintf("URL: %s\n", feed.URL)
	}

	return content
}

func feedRemoveCommand(subscriptionID int) string {
	subscription, err := models.FindSubscriptionG(subscriptionID)
	if subscription == nil {
		return fmt.Sprintf("フィード ID : %d を見つけられませんでした。", subscriptionID)
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

	return fmt.Sprintf("フィード %d (%s) が削除されました", subscription.SubscriptionID, feed.Title)
}
