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

	textMessage := linebot.NewTextMessage(m)
	imageMessage := linebot.NewImageMessage(url, url)
	res, err := bot.PushMessage(l.Events[0].Source.UserID, textMessage, imageMessage).Do()
	if err != nil {
		return err
	}

	slog.Info("Push Message Response", res)

	return nil
}
