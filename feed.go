package main

import (
	"discord-rss-go/models"

	"github.com/volatiletech/sqlboiler/boil"
)

func createFeed(url string, title string) (*models.Feed, error) {
	feed := &models.Feed{
		URL:   url,
		Title: title,
	}
	err := feed.InsertG(boil.Infer())

	return feed, err
}

func getFeeds() (models.FeedSlice, error) {
	feeds, err := models.Feeds().AllG()

	return feeds, err
}

func deleteFeed(feedID int) (*models.Feed, error) {
	feed, err := models.FindFeedG(feedID)
	if feed == nil {
		return nil, err
	}

	_, err = feed.DeleteG()
	if err != nil {
		return nil, err
	}

	return feed, err
}

func existsFeed(url string) (bool, error) {
	ok, err := models.Feeds(
		models.FeedWhere.URL.EQ(url),
	).ExistsG()
	return ok, err
}
