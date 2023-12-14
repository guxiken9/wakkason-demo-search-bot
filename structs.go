package main

import (
	"encoding/json"
)

// Lambda Function URL のリクエスト
type LambdaFunctionURLRequest struct {
	Body string `json:"body"`
}

func UnmarshalWelcome(data []byte) (LambdaFunctionURLRequest, error) {
	var r LambdaFunctionURLRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (l *LambdaFunctionURLRequest) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

// LINE Messaging API でPOSTされるデータ構造
type LineRequestBody struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

func UnmarshalLineRequestBody(data []byte) (LineRequestBody, error) {
	var r LineRequestBody
	err := json.Unmarshal(data, &r)
	return r, err
}

func (l *LineRequestBody) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

type Event struct {
	Type            string          `json:"type"`
	Message         Message         `json:"message"`
	WebhookEventID  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	ReplyToken      string          `json:"replyToken"`
	Mode            string          `json:"mode"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Message struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	QuoteToken string `json:"quoteToken"`
	Text       string `json:"text"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}
