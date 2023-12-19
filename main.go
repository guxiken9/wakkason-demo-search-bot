package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event LambdaFunctionURLRequest) (string, error) {
	slog.Info("Start")
	// JSON のデコード
	l, err := UnmarshalLineRequestBody([]byte(event.Body))
	if err != nil {
		return "", err
	}
	searchWord := l.Events[0].Message.Text

	// TiDBから検索（画像のURL含む）
	result, err := Search(searchWord)
	if err != nil {
		slog.Error("DB Select error ", err)
		return "", err
	}

	for _, r := range result {
		if err := replyToLINE(l, r.Title, r.PhotoURL); err != nil {
			slog.Error("Reply to LINE Error ", err)
			return "", err
		}
	}

	slog.Info("End")
	return "### success ###", nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	lambda.Start(HandleRequest)
}
