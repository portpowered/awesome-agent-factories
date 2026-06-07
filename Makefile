.PHONY: all lint check links fmt test

all: check

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "golangci-lint: not found on PATH (required for Go static analysis)" >&2; \
		echo "Install: go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest" >&2; \
		echo "See also: https://golangci-lint.run/welcome/install/" >&2; \
		exit 1; \
	}
	golangci-lint run --config .golangci.yml ./...

check:
	go run ./internal/checks

links:
	@command -v lychee >/dev/null 2>&1 || { \
		echo "lychee: not found on PATH (required for markdown link checking)" >&2; \
		echo "Install: brew install lychee" >&2; \
		echo "See also: https://lychee.cli.rs/guides/getting-started/" >&2; \
		exit 1; \
	}
	lychee --config .lychee.toml README.md docs/*.md

fmt:
	gofmt -w internal

test:
	go test ./...
