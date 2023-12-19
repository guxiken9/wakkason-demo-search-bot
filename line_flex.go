package main

import (
	"encoding/json"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func UnmarshalFlexMesage(data []byte) (FlexMesage, error) {
	var r FlexMesage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FlexMesage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FlexMesage struct {
	Type string `json:"type"`
	Hero Hero   `json:"hero"`
	Body Body   `json:"body"`
}

type Body struct {
	Type     string    `json:"type"`
	Layout   string    `json:"layout"`
	Contents []Content `json:"contents"`
}

type Content struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Weight string `json:"weight"`
	Wrap   bool   `json:"wrap"`
	Size   string `json:"size"`
}

type Hero struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Size        string `json:"size"`
	AspectRatio string `json:"aspectRatio"`
	AspectMode  string `json:"aspectMode"`
}

func NewFlex(url, text string) (*linebot.FlexMessage, error) {

	m := FlexMesage{
		Type: "bubble",
		Hero: Hero{
			Type:        "image",
			URL:         url,
			Size:        "full",
			AspectRatio: "20:13",
			AspectMode:  "cover",
		},
		Body: Body{
			Type:   "box",
			Layout: "vertical",
			Contents: []Content{
				{
					Type:   "text",
					Text:   text,
					Weight: "regular",
					Wrap:   true,
					Size:   "xl",
				},
			},
		},
	}

	jsonStr, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	container, err := linebot.UnmarshalFlexMessageJSON(jsonStr)
	if err != nil {
		// 正しくUnmarshalできないinvalidなJSONであればerrが返る
		return nil, err
	}
	message := linebot.NewFlexMessage("alt text", container)

	return message, nil

}
