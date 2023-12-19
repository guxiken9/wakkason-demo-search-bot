package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const AWS_REGION = "ap-northeast-1"
const S3_BUCKET_NAME = "wakkason-demo-bucket-team-e"

func NewS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
	})
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)
	return client, nil
}
