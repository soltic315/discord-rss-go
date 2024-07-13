package main

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
)

func sendMessage(channelID string, msg string) error {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		return err
	}
	return nil
}

func feedSubscribeCommand(url string) string {
	var content string

	ok, err := existsFeed(url)
	if err != nil {
		slog.Error("error", err)
		content = "フィードの取得中に不具合が発生しました"
		return content
	}
	if ok {
		content = fmt.Sprintf("このチャンネルにはすでにそのフィード（%s）が登録されています", url)
		return content
	}

	rawFeed, err := fetchRawFeed(url)
	if err != nil {
		slog.Error("error", err)
		content = "フィードの取得中に不具合が発生しました"
		return content
	}
	slog.Debug("Fetch feed", "url", url)

	feed, err := createFeed(url, rawFeed.Title)
	if err != nil {
		slog.Error("error", err)
		content = "フィードの取得中に不具合が発生しました"
		return content
	}

	slog.Info("Create feed", "feedID", feed.FeedID)

	for _, item := range rawFeed.Items {
		if item == nil {
			break
		}

		title := item.Title
		link := item.Link
		publishedAt := *item.PublishedParsed

		article, err := createArticle(feed.FeedID, title, link, publishedAt)
		if err != nil {
			slog.Error("error", err)
			continue
		}
		slog.Debug("Create article", "title", article.Title)
	}

	content += fmt.Sprintf("購読フィード: %s\n", feed.Title)
	content += fmt.Sprintf("[%s](%s)\n", rawFeed.Items[0].Title, rawFeed.Items[0].Link)
	content += rawFeed.Items[0].Description

	return content
}

func feedListCommand() string {
	var content string

	feeds, err := getFeeds()
	if err != nil {
		slog.Error("error", err)
		content = "フィードの取得中に不具合が発生しました"
		return content
	}
	if feeds == nil {
		content = "フィードがありません"
		return content
	}

	for _, feed := range feeds {
		feedID := feed.FeedID
		title := feed.Title
		url := feed.URL

		content += fmt.Sprintf("ID: %d - タイトル: %s\n", feedID, title)
		content += fmt.Sprintf("URL: %s\n", url)
	}

	return content
}

func feedRemoveCommand(feedID int) string {
	var content string

	feed, err := deleteFeed(feedID)
	if err != nil {
		slog.Error("error", err)
		content = "フィードの取得中に不具合が発生しました"
		return content
	}
	if feed == nil {
		content = fmt.Sprintf("フィード ID : %d を見つけられませんでした。", feedID)
		return content
	}
	slog.Info("Delete feed", "feedID", feedID)
	content = fmt.Sprintf("フィード %d (%s) が削除されました", feed.FeedID, feed.Title)
	return content
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		var content string
		options := i.ApplicationCommandData().Options
		if len(options) > 0 {
			switch options[0].Name {
			case "subscribe":
				options = options[0].Options
				url := options[0].StringValue()
				content = feedSubscribeCommand(url)
			case "list":
				content = feedListCommand()
			case "remove":
				options = options[0].Options
				feedID := int(options[0].IntValue())
				content = feedRemoveCommand(feedID)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
		}
	}
}

func createCommand() *discordgo.ApplicationCommand {
	childCommands := []*discordgo.ApplicationCommandOption{
		{
			Name:        "subscribe",
			Description: "このチャンネルでフィードを購読したい場合 :（入力例）/feed subscribe http://kotaku.com/vip.xml",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "url",
					Description: "url",
					Required:    true,
				},
			},
			Type: discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "list",
			Description: "このチャンネルで購読しているフィードを一覧表示する場合 : /feed list",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "remove",
			Description: "このチャンネルからフィードを削除する場合 : /feed remove フィードID",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "feed_id",
					Description: "feed_id",
					Required:    true,
				},
			},
			Type: discordgo.ApplicationCommandOptionSubCommand,
		},
	}
	parentCommand := &discordgo.ApplicationCommand{
		Name:        "feed",
		Description: "有効なコマンド : subscribe、list、remove",
		Options:     childCommands,
	}

	return parentCommand
}
