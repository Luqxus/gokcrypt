build:
	@go build -o ./bin/kcrypt

run: build
	@./bin/kcrypt

test:
	go test ./...

