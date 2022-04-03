.PHONY: build

build:
	go build ./cmd/app/main.go

.DEFAULT_GOAL = build