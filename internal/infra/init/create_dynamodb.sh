#!/bin/sh
set -eu

AWS_ENDPOINT="http://localhost:4566"
AWS_REGION="us-east-1"

export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=${AWS_REGION}

wait_for_service() {

  until aws --endpoint-url "$AWS_ENDPOINT" dynamodb list-tables >/dev/null 2>&1; do
    echo "[localstack-init] Waiting for DynamoDB..."
    sleep 2
  done
}

wait_for_service

aws --endpoint-url "$AWS_ENDPOINT" dynamodb create-table \
  --table-name URL \
  --attribute-definitions AttributeName=Id,AttributeType=S AttributeName=Code,AttributeType=S \
  --key-schema AttributeName=Id,KeyType=HASH AttributeName=Code,KeyType=RANGE \
  --billing-mode PAY_PER_REQUEST \
  --global-secondary-indexes '[{"IndexName":"code-index","KeySchema":[{"AttributeName":"Code","KeyType":"HASH"}],"Projection":{"ProjectionType":"ALL"}}]' >/dev/null

