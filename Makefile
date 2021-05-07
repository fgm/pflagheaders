all: lint

.PHONY: lint
lint:
	golint -min_confidence=0.3 .
	staticcheck .
