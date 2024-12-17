.PHONY: all build install

all:
	go run ./cmd/space/main.go

build:
	go build -o ./bin/space ./cmd/space/main.go

install:
	cd ./cmd/space; go install 
