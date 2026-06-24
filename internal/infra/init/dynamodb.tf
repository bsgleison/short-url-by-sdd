resource "aws_dynamodb_table" "url" {
  name         = "URL"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"
  range_key    = "Code"

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "Code"
    type = "S"
  }

  global_secondary_index {
    name            = "code-index"
    hash_key        = "Code"
    projection_type = "ALL"
  }
}