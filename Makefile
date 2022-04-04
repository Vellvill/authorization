.PHONY: build

build:
	go build ./cmd/app/main.go
.PHONY: test
test:
	go test -v -race -timout 30s ./ ...

.DEFAULT_GOAL = build