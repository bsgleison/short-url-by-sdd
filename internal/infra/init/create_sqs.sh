#!/bin/sh
set -eu

AWS_ENDPOINT="http://localhost:4566"
AWS_REGION="us-east-1"
QUEUE_NAME="url-clicked.fifo"

export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=${AWS_REGION}

wait_for_service() {
  until aws --endpoint-url "$AWS_ENDPOINT" sqs list-queues >/dev/null 2>&1; do
    echo "[localstack-init] Waiting for SQS..."
    sleep 2
  done
}

wait_for_service

aws --endpoint-url "$AWS_ENDPOINT" sqs create-queue \
  --queue-name "$QUEUE_NAME" \
  --attributes '{"FifoQueue":"true","ContentBasedDeduplication":"true"}' >/dev/null

QUEUE_URL=$(aws --endpoint-url "$AWS_ENDPOINT" sqs get-queue-url \
  --queue-name "$QUEUE_NAME" \
  --query 'QueueUrl' \
  --output text)

echo "[localstack-init] Created SQS queue: ${QUEUE_URL}"
