package main

import (
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/go-resty/resty/v2"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

const AWS_REGION = "ap-northeast-1"
const KENDRA_INDEX_ID = "eb17ff24-49b1-4851-8285-9b16b2ffa6a4"

func downloadFile(url string) (*resty.Response, error) {

	c := resty.New()

	// ファイルをダウンロード
	resp, err := c.R().Get(url)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return resp, err
}

func newKendraClient() (*kendra.Kendra, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
	})
	if err != nil {
		return nil, err
	}

	client := kendra.New(sess)
	return client, nil

}

func kendraSearch(keyword string) (*kendra.QueryOutput, error) {

	c, err := newKendraClient()
	if err != nil {
		return nil, err
	}

	// 検索オプション
	attributeFilter := &kendra.AttributeFilter{
		AndAllFilters: []*kendra.AttributeFilter{
			{
				EqualsTo: &kendra.DocumentAttribute{
					Key: aws.String("_language_code"),
					Value: &kendra.DocumentAttributeValue{
						StringValue: aws.String("ja"),
					},
				},
			},
		},
	}

	input := &kendra.QueryInput{
		QueryText:       aws.String(keyword),
		IndexId:         aws.String(KENDRA_INDEX_ID),
		PageNumber:      aws.Int64(1),
		PageSize:        aws.Int64(1),
		AttributeFilter: attributeFilter,
	}
	slog.Info(input.String())

	result, err := c.Query(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func replyToLINE(l LineRequestBody, m, url string) error {

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
	searchWord := l.Events[0].Message.Text

	// Kendra で検索
	result, err := kendraSearch(searchWord)
	if err != nil {
		slog.Error("Kendra Search Error ", err)
		return "", err
	}
	slog.Info(result.String())
	var message string
	var url string
	for _, r := range result.ResultItems {
		url = *r.DocumentURI
		message += *r.DocumentTitle.Text
		message += *r.DocumentExcerpt.Text
	}

	// 画像取得 今はKendraのPDF
	_, err = downloadFile(url)
	if err != nil {
		slog.Error("Download File Error ", err)
		return "", err
	}

	// 検索結果をLINEあてに返却
	if err := replyToLINE(l, message, url); err != nil {
		slog.Error("Reply to LINE Error ", err)
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
