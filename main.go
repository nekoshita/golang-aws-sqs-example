package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/nekoshita/golang-aws-sqs-example/src/infra"
)

type config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	SQSQueueURL        string
	SQSRegion          string
}

func main() {
	conf := getConfig()

	// sqsClient, err := infra.NewSQSClient(
	// 	"https://sqs.ap-northeast-1.amazonaws.com/802904416230/my_queue.fifo",
	// 	"ap-northeast-1",
	// 	"AKIA3V4GVD7TEUX6YQHZ",
	// 	"O8/cnlK7m8iBC27a4uh2vRYanRjkvnrq4ibqMlTl",
	// )
	sqsClient, err := infra.NewSQSClient(
		conf.SQSQueueURL,
		conf.SQSRegion,
		conf.AWSAccessKeyID,
		conf.AWSSecretAccessKey,
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// messageを作成する
	if err := sqsClient.PutMessage(ctx, "sampleMessgaeGroupID", "sameplMessageDuplicateID", "sampleMessage"); err != nil {
		log.Fatal(err)
	}

	// messageを取得する
	sqsMessage, err := sqsClient.GetMessage(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", sqsMessage)

	// messageを削除する
	if err := sqsClient.DeleteMessage(ctx, sqsMessage.ReceiptHandle); err != nil {
		log.Fatal(err)
	}
}

func getConfig() *config {
	// read credentials from environment variables if available
	conf := &config{
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SQSQueueURL:        os.Getenv("SQS_QUEUE_URL"),
		SQSRegion:          os.Getenv("SQS_REGION"),
	}
	// allow consumer credential flags to override confing fields
	awsAccessKeyID := flag.String("aws-access-key-id", "", "AWS Access Key ID")
	awsSecretAccessKey := flag.String("aws-secret-access-key", "", "AWS Access Key ID")
	sqsQueueURL := flag.String("sqs-queue-url", "", "SQS Queue URL")
	sqsRegion := flag.String("sqs-region", "", "SQS Region")
	flag.Parse()
	if *awsAccessKeyID != "" {
		conf.AWSAccessKeyID = *awsAccessKeyID
	}
	if *awsSecretAccessKey != "" {
		conf.AWSSecretAccessKey = *awsSecretAccessKey
	}
	if *sqsQueueURL != "" {
		conf.SQSQueueURL = *sqsQueueURL
	}
	if *sqsRegion != "" {
		conf.SQSRegion = *sqsRegion
	}

	return conf
}
