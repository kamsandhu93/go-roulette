GO_VERSION ?= 1.17.5

run:
	go${GO_VERSION} run .

build:
	go${GO_VERSION} build .

tidy:
	go${GO_VERSION} mod tidy

.PHONY: run