package publisher

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type URLClickedPublisher struct {
	client   *sqs.Client
	queueURL string
}

func NewURLClickedPublisher(client *sqs.Client, queueURL string) *URLClickedPublisher {
	return &URLClickedPublisher{
		client:   client,
		queueURL: queueURL,
	}
}

func (p *URLClickedPublisher) Publish(ctx context.Context, code string) error {
	if p == nil || p.client == nil || p.queueURL == "" {
		return nil
	}

	payload, err := json.Marshal(map[string]string{"code": code})
	if err != nil {
		return err
	}

	_, err = p.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:       aws.String(p.queueURL),
		MessageBody:    aws.String(string(payload)),
		MessageGroupId: aws.String(code),
	})
	return err
}
