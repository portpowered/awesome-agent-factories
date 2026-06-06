.PHONY: all lint check links fmt test

all: check

lint:
	golangci-lint run ./...

check:
	go run ./internal/checks

links:
	lychee README.md docs/*.md

fmt:
	gofmt -w internal

test:
	go test ./...
