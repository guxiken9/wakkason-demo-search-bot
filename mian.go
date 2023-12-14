package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event LambdaFunctionURLRequest) (string, error) {
	fmt.Println("### Start ###")

	// JSON のデコード
	_, err := UnmarshalLineRequestBody([]byte(event.Body))
	if err != nil {
		return "", err
	}

	// LINEのリクエストから検索キーワードを取得

	// 検索キーワードを元に検索を実行

	// 検索結果をLINEあてに返却

	fmt.Println("### End ###")
	return "### success ###", nil
}

func main() {
	lambda.Start(HandleRequest)
}
