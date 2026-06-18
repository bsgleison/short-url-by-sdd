module github.com/bsgleison/short-url-by-sdd

go 1.24

require (
	github.com/aws/aws-sdk-go-v2 v1.42.0
	github.com/aws/aws-sdk-go-v2/config v1.31.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.7.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.35.0
	github.com/go-chi/chi/v5 v5.3.0
	github.com/google/uuid v1.6.0
)

require github.com/aws/aws-sdk-go-v2/service/sqs v1.44.0 // indirect

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.18.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.12.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.33.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.37.0 // indirect
	github.com/aws/smithy-go v1.27.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.5.1
)
