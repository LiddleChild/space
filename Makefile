all:
	go run ./cmd/space/main.go

build:
	go build -o space ./cmd/space/main.go

install:
	cd ./cmd/space; go install
