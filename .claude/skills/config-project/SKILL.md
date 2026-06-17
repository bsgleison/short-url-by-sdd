# SKILL
Name: config-api-project
Friendly name: Config API Project
Dedscription: Configure the basic struture for a new API Go project

## How to execute
- Ask to user the module name
  - module name should respect the format: `example.com\module-name`
- Create the base root layout folders [Root layout]
- Create the internal layout folders [Internal package organization]
- Create the test layout folders [Testing definitions]
- Create basic shared folders and files, use the references in `.\references\internal\shared`
- Create the .env file in root project folder
- Create a `.gitignore` file to go application in root project folder

## Root layout

- `cmd/`: entrypoints for executable applications. 
  - Create a basic 'hello world' as entrypoint in `cmd/api/main.go` for the web API server.
- `internal/`: core implementation packages. This follows Go convention to keep application code private to the module.
- `test/`: test support artifacts and fakes, separate from production code.
- `go.mod`: Go module root with module name asked to user.
- `docker-compose.yml`: local orchestration configuration.

## Internal package organization

- `internal/application/`: application layer that orchestrates use cases.
  - `models/`: request/response or transfer models used by handlers and use cases.
  - `service/`: domain service abstractions and business helpers.
  - `usecase/`: use case implementations and application logic.
- `internal/domain/`: domain model definitions and repository interfaces.
  - `entity/`: domain entities and their behavior.
  - `repository/`: repository interfaces for persistence operations.
- `internal/handler/`: transport layer for HTTP handlers.
  - `http/`: HTTP route handlers and request/response handling logic.
- `internal/infra/`: infrastructure implementations and integration code.
  - `database/`: persistence implementations.
    - `repository/`: concrete database repository packages.
  - `init/`: environment or local startup scripts.
  - `messaging/`: messaging integration.
    - `consumer/`: queue consumer implementations.
    - `publisher/`: queue publisher implementations.
- `internal/shared/`: shared utilities and cross-cutting concerns.
  - `validation/`: validation results and error message types.

## Testing definitions

- `test/faker/` contains test doubles and fake repository implementations for unit tests.
