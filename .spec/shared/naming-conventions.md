# Naming patterns and conventions

Define patterns and conventions for naming files, folders in a Go context application. Includes best practices to work with agentic spec development.

## General rules

- Names should indicate the **reponsability** instead implementation

## Directories

- Directory names are lowercase, singular or conventional compound names: `cmd`, `internal`, `service`, `usecase`, `repository`, `handler`, `publisher`, `consumer`.
- When necessary use `snake_case` style to naming folders.

## Packages

- Go package names are lowercase and match the directory name.

## File name

- File names use lowercase `snake_case` style and describe the package or type, e.g. `url_repository.go`, `click_short_url_usecase.go`, `short_url_handler.go`.

### File Sufix rule

Respect this sufix rules when creating any of these files:

- `*_handler.go` - sufix for all handler files, e.g. `short_url_handler.go`
- `*_repository.go` - sufix for all repositories files, e.g. `short_url_repository.go`
- `*_usecase.go` - sufix for all use cases files, e.g. `create_short_url_usecase.go`
- `*_service.go` - sufix for all services files, e.g. `short_url_service.go`
- `*_test.go` - sufix for all tests files, e.g. `create_short_url_usecase_test.go`
- `*_entity` - sufix for all entities files, e.g. `url_entity.go`
- `*_publisher` - sufix for all queue publisher, e.g. `url_clicked_publisher.go`
- `*_consumer` - sufix for all queue consumers, e.g. `url_clicked_consumer.go`

## Implementation convetions

- Constructors follow `New<TypeName>` naming, e.g. `NewURLRepository`, `NewShortUrlHandler`, `NewClickURLUseCase`.
- Never use technology names when naming constructos, structs, interfaces, e.g. `NewDynamoDBURLRepository`.
- Use `...Request` to all models tha represet parameters receive by API and Use Cases
- Use `...Response` to all models tha represet response returned by API and Use Cases

## Conventions for `Changes`

- are created inside `.spec/changes`
- Sequential prefix with 3 digits
- Format `NNN-change-description/spec.md`
- description should use `kebab_case` style
- the `spec.md` is a market pattern

Example:
- `001-create-project/spec.md`

## Controlled Exceptions

Files like `SKILL.md`, `TODO.md` are allowed.
