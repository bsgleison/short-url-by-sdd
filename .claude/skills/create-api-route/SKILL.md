# SKILL
Name: create-api-route
Friendly name: Create a new API route
Dedscription: Create a new API route int the current project.

## How to execute
- Create the route and use cases basead on details prived by user.
- The use case name, route name, entity name shoud reference what user are requesting
- Implment if needed:
  - Repository
  - Database
  - Entity (with its rules)
  - Use Cases
  - Models
  - Tests
- Ensure all tests are running sueccesfuly
- Use the `.env` file to store dynamic configuration (like AWS Local Stack config, DynamoDB confi....)

### General Conventions

Use the follow orientations to keep a standard while implementing the solution:

- `./memory/structure.md` define a basic Go project structure
- `./shared/how-to-execute.md` define how to execute the spec
- `./shared/naming-conventions.md` rules to naming files and folders
- `./shared/technical-context.md` details about tech stack

All uses case must return a Use Case result with success or error. (use the struct s in `internal/shared/validation`)
In case of error the api should return the details, example:

```json
{
	"has_error": true,
	"has_warning": false,
	"messages": [
		{
			"Code": "VALIDATION_ERROR",
			"Type": 0,
			"Description": "Filed LongURL must be informed"
		}
	]
}
```
