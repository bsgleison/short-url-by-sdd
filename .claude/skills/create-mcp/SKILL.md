# SKILL
Name: create-mcp
Friendly name: Create MCP Skill
Dedscription: Create MCP for the currente project based on docs and reference provided for the user.

## How to execute
- Create a MCP using Node.js
- Request the MCP Name for the user
- Request for the user all details and document refereces for the API that will be integrated in the MCP. (Routes, Input and Output payloads.
- If document and references wasn't enough, ask for more details.
- When MCP is done ask the user if would like to install the MCP created. If yes install the MCP.
- Dont run any test, user will test and validate it.

## Rules and Conventions
- Create the MCP in `mcp/<mcp-name>/SKILL.md`
- The mcp name should respect the `kebab-case` style
