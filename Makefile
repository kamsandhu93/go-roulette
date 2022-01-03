install:
	go get .

run:
	go run .

build:
	go build .

clean:
	rm -f go-roulette

test:
	go test .

tidy:
	go mod tidy

.PHONY: install run build clean test tidy