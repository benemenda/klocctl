all: build run  test-coverage clean

build:
	go get -t -v ./...
	go install
	go build -o klocctl main.go

build-release:
	go build -o $(BIN) -ldflags="-s -w" main.go

test-coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

display-coverage: test-coverage
	go tool cover -html=coverage.txt -o coverage.html

test-report:
	go test -coverprofile=coverage.txt -covermode=atomic -json ./... > test_report.json

run:
	./klocctl -h
clean:
	go clean
