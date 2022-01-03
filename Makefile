GO_VERSION ?= 1.17.5

install:
	go${GO_VERSION} get .

run:
	go${GO_VERSION} run .

build:
	go${GO_VERSION} build .

clean:
	rm -f go-roulette

test:
	go${GO_VERSION} test .

tidy:
	go${GO_VERSION} mod tidy

.PHONY: install run build clean test tidy