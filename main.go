package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
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

	BotToken         = os.Getenv("BOT_TOKEN")
	ChannelID        = os.Getenv("CHANNEL_ID")
	CrawlingInterval = 1
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
		ticker := time.NewTicker(time.Duration(CrawlingInterval) * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			crawl()
		}
	}()
}

func main() {
	// Setup DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("error", err)
		panic(err)
	}
	defer db.Close()
	boil.SetDB(db)

	// Setup BOT
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		slog.Error("error", err)
		panic(err)
	}

	s.AddHandler(commandHandler)

	err = s.Open()
	if err != nil {
		slog.Error("error", err)
		panic(err)
	}
	defer s.Close()

	_, err = s.ApplicationCommandCreate(s.State.User.ID, "", createCommand())
	if err != nil {
		slog.Error("error", err)
		panic(err)
	}

	slog.Info("Bot running")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
