run:
	@air

build:
	@go build -o ./bin/shop ./cmd

test:
	@go test -v ./...

