# CLAUDE.md

## Project Overview

See README.md for user-facing documentation, features, and usage examples.

## Development Notes

- Do NOT do git commits. Let the user do those.
- If you need to run "add" command, ask the user to run it since it is interactive.


## Architecture

DevHttps follows a simple CLI architecture:

- **`main.go`**: Entry point that delegates to `cmd.Execute()`
- **`cmd/`** package: CLI command implementations using `urfave/cli/v3`

- **`internal/`** package: Core business logic modules
  - `certbot/`: Certificate management via certbot CLI (`~/.devhttps/certbot/`)
    - `Certificates()`, `CertificatesVerbose()`: list certs (source of truth for available domains)
    - `Run(domain)`: interactive certbot for new domain
  - `caddy/`: Caddy integration (`~/.devhttps/caddy/`)
  - `validate/`: Input validation
  - `hostcheck/`: DNS and hosts file resolution checking

## Common Development Commands

```bash
# Build the binary
go build -o devhttps

# Run the CLI directly (without building)
go run main.go [command] [args]

# Run specific commands
go run main.go add dev.example.com
go run main.go run dev.example.com 3000
go run main.go show
go run main.go check
```

## Code Organization Patterns

- **Command handling**: Each subcommand is a separate file in `cmd/` that returns a `*cli.Command`
- **Input validation**: All user input (domain, port) goes through `validate/` package before business logic
- **No persistent config**: Certbot certificate store is the source of truth for available domains
- **Error handling**: Commands use `cli.Exit("", 1)` for non-zero exit on errors

## Dependencies

- `github.com/urfave/cli/v3`: CLI framework for command parsing and execution
- Standard library only for all internal logic (no external dependencies for business logic)
