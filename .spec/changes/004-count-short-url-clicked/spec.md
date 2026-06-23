# Change 004 - Count short URL Clicked

## Objective

1. Consume the short url events to count (sum) the Clicks of related url on database. Each event represent a new click tha must be added to Clicks property.

## Technicall Context

The event is a SQS message on queue tha is produced by last change we did (`003-short-url-clicked`)

What consumer should do:

1. Consume each event
2. Query the short url in database to get the current Clicks number
2.1. In case of short url not found, the message should be discarted
3. Use a domain rule to add the new click and last used datetime
4. Update Clicks and LastUsed properties in database with new values
5. In case of error (e.g database unavailable) the consumer should retry (with interval of 5 seconds) until database is available again
6. Before redirect the app should publish a SQS event with the Code of clicked url
  6.1. When publish the event apply a groupId by Code, because these messages will be processed in a sequence way. (the consumer will be implemented in another change)

### General Conventions

Use the follow orientations to keep a standard while implementing the solution:

- `./memory/structure.md` define a basic Go project structure
- `./shared/how-to-execute.md` define how to execute the spec
- `./shared/naming-conventions.md` rules to naming files and folders
- `./shared/technical-context.md` details about tech stack

## Tasks

1. [] Implement the consumer `URLClickedConsumer` with retry capabilities and validation to discart invalid messages.
2. [] Implement the use case `URLClickedUseCase` and use a domain rule to add Clicks to the Entity an update LastUsed property.
3. [] Perssist theses changes in the database.
