.PHONY: clean run test

clean:
	go clean
	go mod tidy
	rm *.out
	rm coverage.html

run:
	. .env
	go get .
	go run main.go

test:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out -v
	go tool cover -html coverage.out -o coverage.html