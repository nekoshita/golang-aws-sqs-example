package infra

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsClientImpl struct {
	queueURL string
	sess     *session.Session
}

func NewSQSClient(
	queueURL string,
	awsRegion string,
	awsAccessKeyID string,
	awsSecretKey string,
) (sqsClientImpl, error) {
	awsConfig := aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretKey, ""),
	}
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return sqsClientImpl{}, fmt.Errorf("failed to create aws session: %w", err)
	}
	return sqsClientImpl{
		queueURL: queueURL,
		sess:     sess,
	}, nil
}
func (s *sqsClientImpl) PutMessage(ctx context.Context, messageGroupID, messageDuplicateID, message string) error {
	svc := sqs.New(s.sess)
	if _, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId:         aws.String(messageGroupID),
		MessageBody:            aws.String(message),
		MessageDeduplicationId: aws.String(messageDuplicateID),
		QueueUrl:               aws.String(s.queueURL),
	}); err != nil {
		return fmt.Errorf("failed to send sqs message: %w", err)
	}
	return nil
}

// GetMessage return a sqs message or nil
func (s *sqsClientImpl) GetMessage(ctx context.Context) (*sqs.Message, error) {
	svc := sqs.New(s.sess)
	msg, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(1),
		QueueUrl:            aws.String(s.queueURL),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to recieve sqs message: %w", err)
	}
	if len(msg.Messages) == 0 {
		return nil, nil
	}
	return msg.Messages[0], nil
}

func (s *sqsClientImpl) DeleteMessage(ctx context.Context, receiptHandle *string) error {
	svc := sqs.New(s.sess)
	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(s.queueURL),
		ReceiptHandle: aws.String(*receiptHandle),
	})
	if err != nil {
		return fmt.Errorf("failed to delete sqs message: %w", err)
	}
	return nil
}
