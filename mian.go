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

	fmt.Println("### End ###")
	return "### success ###", nil
}

func main() {
	lambda.Start(HandleRequest)
}
