.PHONY: build
build:
	go build -v ./cmd/botVK

.PHONY: test
test:
	go test -v -race -timeout 40s ./...


.DEFAULT_GOAL := build	

