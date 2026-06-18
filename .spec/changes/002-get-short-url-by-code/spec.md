# Change 001 - Get short URL

## Objective

Create an route to get all details from a short url code

## Technicall Context

Route: /details/{code}

The api should:

1. Receive short url code as path paramenter
2. Validate if the code was receive and has 7 valid digits (base32)
  2.1. If the format is invalid, return a bad request with detais of error
3. Query all informations on database
  3.1. Id
  3.2. Code
  3.4. OriginalUrl
  3.5. ShortUrl
  3.6. Clicks
  3.7. UsedAt
  3.8. CreateAt
4. Use the index to filter by code to improve performance
5. In case of success the api should return all data of short URL.

### General Conventions

Use the follow orientations to keep a standard while implementing the solution:

- `./memory/structure.md` define a basic Go project structure
- `./shared/how-to-execute.md` define how to execute the spec
- `./shared/naming-conventions.md` rules to naming files and folders
- `./shared/technical-context.md` details about tech stack

## Tasks

1. [x] Create the new route using the skill `.claude/skills/create-api-route.SKILL.md`
  > 2026-06-18 00:00 - Added `GET /details/{code}` route, implemented the lookup use case and HTTP handler, validated the code format, used the DynamoDB `code-index` query path, and returned not-found responses for missing entries.
