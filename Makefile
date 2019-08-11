all: build run clean

build:
	go install
	go build -o klocctl main.go
run:
	./klocctl -h
clean:
	go clean
