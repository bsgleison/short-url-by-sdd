resource "aws_sqs_queue" "url_clicked" {
  name                        = "url-clicked.fifo"
  fifo_queue                  = true
  content_based_deduplication = false
}

output "url_clicked_queue_url" {
  value = aws_sqs_queue.url_clicked.url
}
