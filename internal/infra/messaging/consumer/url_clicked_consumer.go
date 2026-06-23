package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
)

type URLClickedConsumer struct {
	client        *sqs.Client
	queueURL      string
	useCase       *usecase.URLClickedUseCase
	retryInterval time.Duration
}

func NewURLClickedConsumer(
	client *sqs.Client,
	queueURL string,
	useCase *usecase.URLClickedUseCase,
) *URLClickedConsumer {
	return &URLClickedConsumer{
		client:        client,
		queueURL:      queueURL,
		useCase:       useCase,
		retryInterval: 5 * time.Second,
	}
}

func (c *URLClickedConsumer) Start(ctx context.Context) {
	log.Println("Starting URL clicked consumer")

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("URL clicked consumer stopping")
			default:
				err := c.processMessage(ctx)
				if err != nil {
					log.Printf("Error processing message: %v", err)
					time.Sleep(c.retryInterval)
				}
			}
		}
	}()
}

func (c *URLClickedConsumer) processMessage(ctx context.Context) error {
	output, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &c.queueURL,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     20,
		VisibilityTimeout:   30,
	})
	if err != nil {
		return err
	}

	if len(output.Messages) == 0 {
		return nil
	}

	msg := output.Messages[0]
	delInput := &sqs.DeleteMessageInput{
		QueueUrl:      &c.queueURL,
		ReceiptHandle: msg.ReceiptHandle,
	}

	var payload models.URLClickedMessage
	if err := json.Unmarshal([]byte(*msg.Body), &payload); err != nil {
		log.Printf("Invalid message format: %v", err)
		_, _ = c.client.DeleteMessage(ctx, delInput)
		return nil
	}

	if payload.Code == "" {
		log.Println("Invalid message: empty code")
		_, _ = c.client.DeleteMessage(ctx, delInput)
		return nil
	}

	if err := c.useCase.Execute(ctx, &payload); err != nil {
		log.Printf("Use case error: %v", err)
		return err
	}

	_, err = c.client.DeleteMessage(ctx, delInput)
	return err
}
