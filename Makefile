all: build run clean

build:
	go get -t -v ./...
	go install
	go build -o klocctl main.go
run:
	./klocctl -h
clean:
	go clean
