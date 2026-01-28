# AGENTS.md

Guidance for agentic coding tools working in this repo.

## Project snapshot
- Go module: github.com/pauslik/gosmapi
- Go version: 1.25.4 (see go.mod)
- API client for StorageMAP; tests talk to live server.
- README is WIP; rely on code/tests for behavior.

## Build commands
- Build all packages: `go build ./...`
- Build current package: `go build`
- Build with race (if needed): `go test -c -race ./...`

## Test commands
- Run all tests: `go test ./...`
- Run tests in package: `go test ./tests`
- Run a single test by name: `go test ./tests -run TestNewClient`
- Run a single test across repo: `go test ./... -run TestStructs`
- Run internal config tests: `go test ./internal/config -run TestFunctions`
- Verbose output: `go test -v ./tests`

## Test prerequisites
- Some tests call a live StorageMAP core API.
- Provide env vars `CORE_ADDRESS` and `ADMIN_TOKEN`.
- `tests/client_test.go` loads `.env`; place it at repo root.
- `tests/servers_test.go` uses `godotenv.Load()` with default `.env`.
- If env vars are missing, tests may `t.Skip` or fail.
- Avoid running live-server tests in CI unless configured.

## Linting / static analysis
- Go vet: `go vet ./...`
- Staticcheck (if installed): `staticcheck ./...`
- Go fmt diff check: `gofmt -w $(rg --files -g '*.go')`
- Goimports (if installed): `goimports -w <files>`

## Formatting
- Use `gofmt` for all Go source files.
- Run gofmt on touched files before commit.
- Keep formatting stable; avoid manual alignment.

## Repository structure
- Top-level Go files implement API resources and client logic.
- `internal/config/` holds config helpers and tests.
- `tests/` holds integration tests against StorageMAP.

## Import style
- Use standard Go grouping: stdlib, blank line, third-party, blank line, local.
- Let `gofmt`/`goimports` manage ordering.
- Avoid unused imports; build will fail with `go test`.

## Naming conventions
- Exported identifiers use `CamelCase`.
- Unexported identifiers use `lowerCamelCase`.
- Types: `Client`, `Config`, `SMuser`, `SMserver`.
- Const enums use all-caps values (`ADMIN`, `CORE`).
- Prefer descriptive receiver names (`c *Client`).

## Types and data modeling
- Prefer concrete structs over `map[string]interface{}`.
- Keep JSON shapes as explicit request/response types.
- Use typed enums (see `userRole`, `serverRole`).
- Avoid global state; use struct fields to carry config.

## Error handling
- Return errors rather than panicking.
- Wrap context with `fmt.Errorf("...: %w", err)` when helpful.
- For HTTP responses, check status codes before unmarshalling.
- Propagate errors from `json.Marshal`/`json.Unmarshal`.
- Use `t.Fatalf` for fatal test errors, `t.Error` for non-fatal.

## Context usage
- `Client.makeRequest` expects a `context.Context`.
- New API methods should accept a `context.Context` parameter.
- Callers can pass `context.Background()` when no cancellation needed.

## HTTP conventions
- Use `c.HTTPClient` (do not create new clients per request).
- Set `Content-Type` to `application/vnd.api+json`.
- Include bearer token when `Client.Token` is set.
- Follow existing pattern for BaseURL: `http://<address>/api`.

## JSON handling
- Use `encoding/json` and explicit struct tags.
- Keep response structs aligned with API payloads.
- Do not ignore JSON errors unless intentionally safe.

## Testing guidelines
- Keep integration tests under `tests/`.
- Keep unit tests close to code (`*_test.go`).
- If a test requires env vars, skip with a clear message.
- Avoid destructive API calls in tests unless documented.

## Documentation updates
- README is minimal; update only when asked.
- Prefer inline doc comments for exported types/functions.

## Adding new packages
- Place internal-only helpers under `internal/`.
- Keep public API surface in root `gosmapi` package.

## Versioning
- `go.mod` targets Go 1.25.4; avoid features requiring newer Go.
- Update `go.sum` via `go mod tidy` if dependencies change.

## Dependency management
- Use `go mod tidy` after adding/removing deps.
- Prefer standard library over new dependencies.
- Document new deps in PR description if asked.

## Compatibility
- Keep public API stable; avoid breaking exported fields.
- If renaming exported identifiers, provide migration notes.

## Security & secrets
- Never commit `.env` files, tokens, or credentials.
- Tests read `ADMIN_TOKEN`; use local env only.

## Known rules sources
- No `.cursor/rules/`, `.cursorrules`, or `.github/copilot-instructions.md` found in repo.
- If new agent rules are added later, follow them.

## Suggested workflows
- Typical change: edit code → `gofmt` → `go test ./...`.
- For integration changes: run `go test ./tests -run TestGetServers`.
- For config changes: run `go test ./internal/config`.

## Single-test examples
- `go test ./tests -run TestNewClient`
- `go test ./tests -run TestAddFileServerGeneric`
- `go test ./internal/config -run TestStructs`
- `go test ./... -run TestFunctions`

## Performance notes
- Avoid per-call allocations when building request payloads.
- Reuse structs when possible in tight loops.

## Code review checklist
- Go fmt applied.
- Errors handled and surfaced.
- Tests updated or justified.
- Public API changes documented.

## When in doubt
- Follow standard Go conventions.
- Match existing patterns in `client.go` and resource files.
- Keep changes minimal and focused.

## End
