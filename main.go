package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		fmt.Println("error creating session: %v", err)
	}

	bucketInput := s3.ListBucketsInput{}

	svc := s3.New(sess)

	buckets, err := svc.ListBuckets(&bucketInput)

	fmt.Printf("%+v", buckets)
}
