package main

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kendra"
)

const KENDRA_INDEX_ID = "eb17ff24-49b1-4851-8285-9b16b2ffa6a4"

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

func KendraSearch(keyword string) (*kendra.QueryOutput, error) {

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
