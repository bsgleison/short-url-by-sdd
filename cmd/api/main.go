package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
	handler "github.com/bsgleison/short-url-by-sdd/internal/handler/http"
	repo "github.com/bsgleison/short-url-by-sdd/internal/infra/database/repository"
	"github.com/bsgleison/short-url-by-sdd/internal/infra/messaging/consumer"
	publisher "github.com/bsgleison/short-url-by-sdd/internal/infra/messaging/publisher"
)

func main() {
	context := context.Background()

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	cfg, err := config.LoadDefaultConfig(context)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config: %v", err)
	}

	log.Println(os.Getenv("AWS_ENDPOINT_URL_DYNAMODB"))

	dynamodbOpts := []func(*dynamodb.Options){}
	sqsOpts := []func(*sqs.Options){}
	if endpoint := os.Getenv("AWS_ENDPOINT_URL_DYNAMODB"); endpoint != "" {
		log.Printf("using custom AWS endpoint: %s", endpoint)
		dynamodbOpts = append(dynamodbOpts, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
		sqsOpts = append(sqsOpts, func(o *sqs.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
	}

	dbClient := dynamodb.NewFromConfig(cfg, dynamodbOpts...)
	sqsClient := sqs.NewFromConfig(cfg, sqsOpts...)

	r := chi.NewRouter()
	urlRepository := repo.NewURLRepository(dbClient)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://short.com"
	}

	newUrlClicUseCase := usecase.NewURLClickedUseCase(urlRepository)

	queueURL := os.Getenv("SQS_QUEUE_URL")
	newUrlClickedPublisher := publisher.NewURLClickedPublisher(sqsClient, queueURL)
	newUrlClickedConsumer := consumer.NewURLClickedConsumer(sqsClient, queueURL, newUrlClicUseCase)
	newUrlClickedConsumer.Start(context)

	createShortURLUseCase := usecase.NewCreateShortURLUseCase(urlRepository, baseURL)
	shortURLHandler := handler.NewCreateShortURLHandler(createShortURLUseCase)
	getShortURLByCodeUseCase := usecase.NewGetShortURLByCodeUseCase(urlRepository)
	getShortURLByCodeHandler := handler.NewGetShortURLByCodeHandler(getShortURLByCodeUseCase)
	redirectShortURLHandler := handler.NewRedirectShortURLHandler(getShortURLByCodeUseCase, newUrlClickedPublisher)

	r.Post("/shorten", shortURLHandler.Create)
	r.Get("/details/{code}", getShortURLByCodeHandler.Get)
	r.Get("/{code}", redirectShortURLHandler.Redirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
