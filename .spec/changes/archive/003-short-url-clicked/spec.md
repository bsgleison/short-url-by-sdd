# Change 003 - Short URL Clicked

## Objective

1. Create an route that redirect user to original URL when the short URL is clicked
2. Update Clicks count to the short url related, this process must be done by a async way

## Technicall Context

Route: /{code}

The api should:

1. Receive short url code as path paramenter
2. Validate if the code was receive and has 7 valid digits (base32)
  2.1. If the format is invalid, return a bad request with detais of error
3. Query the OriginalUrl on database by Code received
4. Use the index to filter by code to improve performance
5. In case of success the api should redirect (http status 301) user to orinigal url
6. Before redirect the app should publish a SQS event with the Code of clicked url
  6.1. When publish the event apply a groupId by Code, because these messages will be processed in a sequence way. (the consumer will be implemented in another change)

### General Conventions

Use the follow orientations to keep a standard while implementing the solution:

- `./memory/structure.md` define a basic Go project structure
- `./shared/how-to-execute.md` define how to execute the spec
- `./shared/naming-conventions.md` rules to naming files and folders
- `./shared/technical-context.md` details about tech stack

## Tasks

1. [x] Create the new route using the skill `.claude/skills/create-api-route.SKILL.md`
   > 2026-06-18 00:00 - Added `GET /{code}` redirect handling and wired it into the API entrypoint.
2. [x] Validate and find the short url, use example of the already implementation in `./internal/usecase/get_short_url_by_code.go`
   > 2026-06-18 00:00 - Reused the existing code validation flow and lookup behavior to resolve the target URL before redirecting.
3. [x] Create the publisher to pusblish the click event
   > 2026-06-18 00:00 - Implemented an SQS publisher that sends the clicked code with a message group ID derived from the code.
