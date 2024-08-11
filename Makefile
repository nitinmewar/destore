run: build
	./bin/destore

build:
	go build -o bin/destore main.go

test:
	go test -v ./...
