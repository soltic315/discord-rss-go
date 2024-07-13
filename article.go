package main

import (
	"discord-rss-go/models"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

func createArticle(feedID int, title string, link string, publishedAt time.Time) (*models.Article, error) {
	article := &models.Article{
		FeedID:      feedID,
		Title:       title,
		Link:        link,
		PublishedAt: publishedAt,
	}
	err := article.InsertG(boil.Infer())

	return article, err
}

func existsArticle(feedID int, publishedAt time.Time) (bool, error) {
	ok, err := models.Articles(
		models.ArticleWhere.FeedID.EQ(feedID),
		models.ArticleWhere.PublishedAt.EQ(publishedAt),
	).ExistsG()
	return ok, err
}
