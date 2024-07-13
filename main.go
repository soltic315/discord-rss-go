package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	DebugMode = flag.Bool("debug", false, "Enable debug mode")

	DBHost     = os.Getenv("DB_HOST")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
	DBPort     = os.Getenv("DB_PORT")
	SslMode    = os.Getenv("SSL_MODE")

	BotToken                 = os.Getenv("BOT_TOKEN")
	CrawlingIntervalMinute   = 5
	CrawlingStopFailureCount = 12
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	// Setup Logger
	level := slog.LevelInfo
	if *DebugMode {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	// Setup Scheduled Jobs
	go func() {
		ticker := time.NewTicker(time.Duration(CrawlingIntervalMinute) * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			crawlingJob()
			notificationJob()
		}
	}()
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		var content string
		userName := i.Member.User.Username

		options := i.ApplicationCommandData().Options
		if len(options) > 0 {
			command := options[0].Name

			switch command {
			case "subscribe":
				options = options[0].Options
				slog.Info("Run Command", "command", command, "userName", userName, "option", options[0].StringValue())
				content = feedSubscribeCommand(options[0].StringValue(), i.ChannelID)
			case "list":
				slog.Info("Run Command", "command", command, "userName", userName)
				content = feedListCommand(i.ChannelID)
			case "remove":
				options = options[0].Options
				slog.Info("Run Command", "command", command, "userName", userName, "option", options[0].IntValue())
				content = feedRemoveCommand(int(options[0].IntValue()))
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   discordgo.MessageFlagsEphemeral,
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

// For Koyeb
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// Setup DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tokyo",
		DBHost, DBUser, DBPassword, DBName, DBPort, SslMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		panic(err)
	}
	defer db.Close()
	boil.SetDB(db)
	boil.DebugMode = *DebugMode

	// Setup Bot
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		slog.Error("Error occurred", "error", err)
		panic(err)
	}

	s.AddHandler(commandHandler)

	err = s.Open()
	if err != nil {
		slog.Error("Error occurred", "error", err)
		panic(err)
	}
	defer s.Close()

	_, err = s.ApplicationCommandCreate(s.State.User.ID, "", createCommand())
	if err != nil {
		slog.Error("Error occurred", "error", err)
		panic(err)
	}

	// For Koyeb
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":8000", nil)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
