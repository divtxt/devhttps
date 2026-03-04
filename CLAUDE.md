# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

DevHttps is a CLI tool for HTTPS reverse proxy in local development. See README.md for user-facing documentation, features, and usage examples.

## Development Notes

- Do NOT do git commits. Let the user do those.
- If you need to run "add" command, ask the user to run it since it is interactive.


## Architecture

DevHttps follows a simple CLI architecture:

- **`main.go`**: Entry point that delegates to `cmd.Execute()`
- **`cmd/`** package: CLI command implementations using `urfave/cli/v3`
  - `root.go`: CLI app setup with command registration
  - `add.go`: Add/update domain-to-port mappings
  - `show.go`: Display configured domain mappings
  - `check.go`: Verify required tools (certbot, caddy) are installed and meet version requirements

- **`internal/`** package: Core business logic modules
  - `config/`: Persistent configuration management (JSON file at `~/.devhttps/config.json`)
    - `Config` struct holds list of domain-to-port entries
    - `Load()`, `Save()`, `Add()` for CRUD operations
  - `validate/`: Input validation
    - `Domain()`: FQDN validation using regex
    - `Port()`: Port number validation (1-65535)
  - `hostcheck/`: DNS and hosts file resolution checking
    - `CheckResolvesToLocalhost()`: Verifies domain resolves to 127.0.0.1 or ::1

## Common Development Commands

```bash
# Build the binary
go build -o devhttps

# Run the CLI directly (without building)
go run main.go [command] [args]

# Run specific commands
go run main.go add dev.myapp.com 3000
go run main.go show
go run main.go check
```

## Code Organization Patterns

- **Command handling**: Each subcommand is a separate file in `cmd/` that returns a `*cli.Command`
- **Input validation**: All user input (domain, port) goes through `validate/` package before business logic
- **Configuration**: All persistent state stored in `~/.devhttps/config.json` via `config` package
- **Error handling**: Commands use `cli.Exit("", 1)` for non-zero exit on errors

## Dependencies

- `github.com/urfave/cli/v3`: CLI framework for command parsing and execution
- Standard library only for all internal logic (no external dependencies for business logic)
