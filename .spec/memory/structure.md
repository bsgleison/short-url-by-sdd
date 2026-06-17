# Project Structure Memory

## Root layout

- `cmd/`: entrypoints for executable applications. Current project has `cmd/api/main.go` for the web API server.
- `internal/`: core implementation packages. This follows Go convention to keep application code private to the module.
- `test/`: test support artifacts and fakes, separate from production code.
- `.spec/`: metadata and specifications directory.
  - `memory/`: agent-oriented structure and project memory (e.g., `structure.md`).
  - `shared/`: shared documentation and context files for team reference.
    - `how-to-execute.md`: execution instructions.
    - `naming-conventions.md`: project-wide naming rules.
    - `technical-context.md`: technical context and stack information.
  - `changes/`: changelogs and breaking changes documentation.
- `go.mod`: Go module root with module name.
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

## Layers Integration definitions

- Use a monorepo structure
- Use a based Clean Archtecture struture
- Domain packages separate entity definitions from repository interface definitions.
- Application layer is responsible for orchestrating use cases and interacting with domain repositories and services.
- Handlers only convert HTTP requests to application models and return HTTP responses.
- Infrastructure packages implement the concrete persistence and messaging behavior used by the application.

## Testing definitions

- `test/faker/` contains test doubles and fake repository implementations for unit tests.
- Add tests alongside packages using `_test.go` naming and place fakes in `test/faker` when appropriate.
- Apply unit test to all domain and use cases

## Agent guidance rules

- Treat `cmd/` as the only executable entrypoint location.
- Do not place production business logic outside `internal/`.
- Keep HTTP transport details in `internal/handler/http`.
- Keep persistence implementations in `internal/infra/database/repository`.
- Keep messaging publisher/consumer implementations in `internal/infra/messaging`.
- Keep reusable validation and shared utilities under `internal/shared`.
- When adding new features, mirror the existing layered structure: domain -> application -> handler -> infra.
