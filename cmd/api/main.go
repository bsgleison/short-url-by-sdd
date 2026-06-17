package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
	handler "github.com/bsgleison/short-url-by-sdd/internal/handler/http"
	repo "github.com/bsgleison/short-url-by-sdd/internal/infra/database/repository"
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

	opts := []func(*dynamodb.Options){}
	if endpoint := os.Getenv("AWS_ENDPOINT_URL_DYNAMODB"); endpoint != "" {
		log.Printf("using custom DynamoDB endpoint: %s", endpoint)
		opts = append(opts, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
	}

	dbClient := dynamodb.NewFromConfig(cfg, opts...)

	r := chi.NewRouter()
	urlRepository := repo.NewURLRepository(dbClient)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://short.com"
	}

	createShortURLUseCase := usecase.NewCreateShortURLUseCase(urlRepository, baseURL)
	shortURLHandler := handler.NewCreateShortURLHandler(createShortURLUseCase)

	r.Post("/shorten", shortURLHandler.Create)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
