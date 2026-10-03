# akash-golang-learning

[![ci](https://github.com/akashgagda/akash-golang-learning/actions/workflows/ci.yml/badge.svg)](https://github.com/akashgagda/akash-golang-learning/actions/workflows/ci.yml)

Akash's Go learning repository for the 20-week run to backend-engineering job
applications (2026-10-03 → 2027-02-19).

## What this is

One repo, built up commit by commit, with a real test suite and a CI pipeline
from day one. A green build here means: formatted, vetted, race-clean, lint-clean.

## What CI checks

`.github/workflows/ci.yml` runs on every push and pull request:

| Step | Command |
|---|---|
| formatting | `gofmt -l .` (fails if anything is unformatted) |
| vet | `go vet ./...` |
| lint | `golangci-lint run` (v2.14.0) |
| tests | `go test -race -count=1 ./...` |

Run the same four locally before pushing:

```bash
gofmt -l . && go vet ./... && golangci-lint run && go test -race -count=1 ./...
```

## Status

`hello.go` is **scaffolding**, placed so CI had something to prove itself
against before the first real commit. Replace it with your own work on Day 1.
