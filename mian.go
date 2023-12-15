package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func replyToLINE(l LineRequestBody, m string) error {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		return err
	}

	token := l.Events[0].ReplyToken
	if _, err := bot.ReplyMessage(token, linebot.NewTextMessage(m)).Do(); err != nil {
		return err
	}
	return nil
}

func HandleRequest(event LambdaFunctionURLRequest) (string, error) {
	slog.Info("Start")
	// JSON のデコード
	l, err := UnmarshalLineRequestBody([]byte(event.Body))
	if err != nil {
		return "", err
	}

	// LINEのリクエストから検索キーワードを取得

	// 検索キーワードを元に検索を実行

	// 検索結果をLINEあてに返却
	if err := replyToLINE(l, "TEST"); err != nil {
		slog.Error("Reply to LINE Error", err)
		return "", err
	}

	slog.Info("End")
	return "### success ###", nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	lambda.Start(HandleRequest)
}
