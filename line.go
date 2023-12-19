package main

import (
	"log/slog"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func replyToLINE(l LineRequestBody, m, url string) error {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return err
	}

	slog.Info(l.Events[0].Source.UserID)
	f, err := NewFlex(url, m)
	if err != nil {
		return err
	}
	_, err = bot.PushMessage(l.Events[0].Source.UserID, f).Do()
	if err != nil {
		return err
	}

	return nil
}
