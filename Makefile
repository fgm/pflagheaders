all: lint

.PHONY: lint
lint:
	gofmt -d -s .
	golint -min_confidence=0.3 .
	staticcheck .
