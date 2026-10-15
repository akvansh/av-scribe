# Contributing to Scribe

Thanks for your interest in contributing! This document covers the workflow and conventions.

## Prerequisites

- Go 1.23+
- [Task](https://taskfile.dev) — task runner
- [Lefthook](https://github.com/evilmartians/lefthook) — git hooks
- Node — only for prettier via `npx` (markdown/YAML formatting), no JS project setup needed

## Getting started

```sh
git clone https://github.com/akvansh/av-scribe.git
cd scribe
task hooks        # install git hooks
task --list       # see available tasks
```

> Note: the Go module is initialized by the maintainer. Until `go.mod` exists, Go tasks
> (`test`, `vet`, `lint`) will no-op or fail — tooling tasks (`fmt:check`) work immediately.

## Workflow

1. Fork and create a feature branch from `main`
2. Make your change, keeping commits scoped to one concern
3. Run `task check` (or at minimum `task fmt:check`)
4. Open a pull request with a clear description of the problem and solution

## Conventions

- **Commits**: [Conventional Commits](https://www.conventionalcommits.org/) —
  `feat:`, `fix:`, `docs:`, `chore:`, `test:`, `perf:`
- **Go**: `gofmt` formatted, stdlib only unless a dependency is well justified,
  table-driven tests, minimal comments (code should explain itself)
- **API**: stays slog-compatible at the core (`Info(msg, attrs...)`) with typed field
  constructors (`Str`, `Dur`, `Err`) inspired by zap
- **Markdown/YAML/JSON**: prettier formatted — hooks do this automatically on commit
- **Scope**: one concern per PR; open an issue first for large changes (new handlers,
  dependencies, API breaks)

## Testing

```sh
task test   # go test -race -coverprofile=coverage.out ./...
```

New behavior needs tests; bug fixes need a regression test.

## Reporting bugs / security

- Bugs: [GitHub issues](https://github.com/akvansh/av-scribe/issues)
- Security: see [security.md](security.md) — please do **not** open a public issue

## Code of conduct

This project follows the [Contributor Covenant](CODE_OF_CONDUCT.md).
