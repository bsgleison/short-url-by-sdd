# Global technical context

**Namespace**: github.com/bsgleison/<module-name-informed-by-user>

## Base stack

- Go is the main language
- Use Docker and Local Stack to run AWS dependencies
- use `.env` file to store e load dynamic configurations
- Dependencies (when need):
  - DynamoDB
  - SQS
- Create a shell script file to create all dependecies on local stack.
  - Save de scripts in `internal/infra/init`
  - Create one script file for each dependecy

## Use

Libraries that you can use (if need):

- `github.com/go-chi/chi/v5`
- `github.com/google/uuid`
- `github.com/joho/godotenv`

### Never use

Never use these libraries:

- `github.com/stretchr/testify`
