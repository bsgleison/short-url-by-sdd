package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
)

type URLRepository struct {
	client *dynamodb.Client
	table  string
}

func NewURLRepository(client *dynamodb.Client) *URLRepository {

	return &URLRepository{
		client: client,
		table:  "URL",
	}
}

func (r *URLRepository) Save(ctx context.Context, url *entity.URL) error {
	if url == nil {
		return nil
	}

	item, err := attributevalue.MarshalMap(urlRecordFromEntity(url))
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      item,
	})
	return err
}

func (r *URLRepository) FindByCode(ctx context.Context, code string) (*entity.URL, error) {
	if code == "" {
		return nil, nil
	}

	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.table),
		FilterExpression: aws.String("Code = :code"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":code": &types.AttributeValueMemberS{Value: code},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	var record urlRecord
	if err := attributevalue.UnmarshalMap(result.Items[0], &record); err != nil {
		return nil, err
	}

	return record.toEntity(), nil
}

type urlRecord struct {
	Id          string    `dynamodbav:"Id"`
	Code        string    `dynamodbav:"Code"`
	OriginalUrl string    `dynamodbav:"OriginalUrl"`
	ShortUrl    string    `dynamodbav:"ShortUrl"`
	Clicks      int       `dynamodbav:"Clicks"`
	UsedAt      time.Time `dynamodbav:"UsedAt"`
	CreateAt    time.Time `dynamodbav:"CreateAt"`
}

func urlRecordFromEntity(url *entity.URL) urlRecord {
	return urlRecord{
		Id:          url.ID,
		Code:        url.Code,
		OriginalUrl: url.OriginalURL,
		ShortUrl:    url.ShortURL,
		Clicks:      url.Clicks,
		UsedAt:      url.UsedAt,
		CreateAt:    url.CreatedAt,
	}
}

func (r urlRecord) toEntity() *entity.URL {
	return &entity.URL{
		ID:          r.Id,
		Code:        r.Code,
		OriginalURL: r.OriginalUrl,
		ShortURL:    r.ShortUrl,
		Clicks:      r.Clicks,
		UsedAt:      r.UsedAt,
		CreatedAt:   r.CreateAt,
	}
}
