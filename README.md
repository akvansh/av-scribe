# Scribe

> Fast, zero-dependency structured logging for Go.

**Status: foundation.** The repo is being scaffolded; the Go module and public API are not
published yet. Track progress in the roadmap below.

## Why Scribe

Scribe is a structured logging library for Go, built around a small handler core you can compose
and extend. It ships the ergonomics you'd expect from zap or zerolog — typed fields, a pretty
console, runtime level switching, and request-scoped loggers — with zero third-party dependencies
and first-class interoperability with the standard library's `log/slog`.

- **Extensible handlers** — a clean handler core; console, NDJSON, or bring your own
- **Typed fields** — `Str`, `Int`, `Dur`, `Err` instead of bare `any` pairs
- **Pretty console** — colored, aligned output for development; NDJSON for production
- **Runtime level switching** — change levels without rebuilding the logger tree
- **Context-aware** — `NewContext` / `FromContext` for request-scoped loggers
- **slog interop** — bridge to and from `log/slog` so the wider ecosystem composes
- **Zero dependencies** — stdlib only

## Install

```sh
go get github.com/akvansh/av-scribe
```

_Not available yet — module init is pending._

## Quickstart (planned)

```go
log := scribe.New(scribe.Pretty(), scribe.WithLevel(scribe.Debug))
log.Info("server started", scribe.Int("port", 8080))

log = log.With(scribe.Str("req_id", rid))
log.Error("request failed", scribe.Err(err))
```

## Structure

```
scribe/
├── scribe.go          # Logger type, default logger, package-level sugar
├── level.go           # Level type, parsing, text marshaling
├── fields.go          # typed field constructors: Str, Int, Dur, Err, ...
├── options.go         # functional options: WithLevel, Pretty, JSON, WithOutput
├── handler.go         # runtime level handler + console handler
├── context.go         # NewContext / FromContext
├── *_test.go          # table-driven tests next to the code
├── cmd/scribe/main.go  # CLI entrypoint (orchestrator / root command)
├── cmd/scribe/cli/     # CLI subcommands (report, inspect, ...)
├── internal/console/   # terminal color/format helpers (private)
├── examples/basic/    # runnable examples, built in CI
└── .github/workflows/ # CI: format + test
```

Each file maps to one concern; the root package stays importable as a single unit
(`import "github.com/akvansh/av-scribe"`).

### Examples

`examples/` contains runnable programs demonstrating real usage — server startup, context
propagation, JSON logging. They compile in CI so documentation never rots. Start with
`examples/basic`.

## Tooling

| Command            | What it does                              |
| ------------------ | ----------------------------------------- |
| `task --list`      | List available tasks                      |
| `task fmt`         | Format Go + markdown/YAML                 |
| `task test`        | Tests with race detector + coverage       |
| `task check`       | All checks (fmt, vet, lint, test)         |
| `lefthook install` | Enable git hooks (format + vet on commit) |

## Roadmap

- [x] Repo scaffolding, tooling, CI
- [ ] Core logger (`New`, levels, typed fields, options)
- [ ] Console + JSON handlers, runtime level switching
- [ ] Context propagation + slog interop
- [ ] Examples, benchmarks, docs
- [ ] v0.1 release
- [ ] Phase 2: `scribe` CLI — log reports and inspection

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). By participating you agree to the
[Code of Conduct](CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE.md)
