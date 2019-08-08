all: build run

build:
	go build -o klocctl main.go
run:
	./klocctl --config klocctl.yaml config print
