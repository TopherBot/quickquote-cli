# quickquote‑cli

A **single‑file** Go CLI that prints a random inspirational quote.

```bash
$ go install github.com/topherbot/quickquote-cli@latest
$ quickquote
"The only limit to our realization of tomorrow is our doubts of today." – Franklin D. Roosevelt
```

## Features
- Zero‑dependency fetch – uses the standard library HTTP client.
- Proper error handling and exit‑codes for CI friendliness.
- Minimal but useful CI with linting, static analysis, and tests.

## Install
```bash
# With Go modules (requires Go 1.23+)
go install github.com/topherbot/quickquote-cli@latest
```

## Run
```bash
quickquote
```

## CI status
[![CI](https://github.com/topherbot/quickquote-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/topherbot/quickquote-cli/actions/workflows/ci.yml)

## License
MIT – see `LICENSE` file.
