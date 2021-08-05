all: lint

.PHONY: lint
lint:
	gofmt -d -s .
	find . -name "*.go" | xargs goimports -d -e
	staticcheck .
