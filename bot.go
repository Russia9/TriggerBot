package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func InitBot(token string, logger *logrus.Logger) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		sentry.CaptureException(err)
		logger.Fatal(err)
		return
	}

	bot.AddHandler(messageCreate)
	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = bot.Open()
	if err != nil {
		sentry.CaptureException(err)
		logger.Fatal(err)
		return
	}

	logger.Info("Bot started successfully")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = bot.Close()
	if err != nil {
		sentry.CaptureException(err)
		logger.Fatal(err)
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
