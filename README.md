Backend API for NPC generation

# Set up instructions:

0. `go mod tidy` to install dependencies
1. Create a `config.env` in the root directory with variables for those listed in `internal/config/config.go`
2. Get local postgres instance running
3. `go run .` and navigate to `localhost:8000` for list of commands

# Todo:

- Replace local postgres with Docker container
- CI pipeline with Travis
- Deploy this somewhere...
