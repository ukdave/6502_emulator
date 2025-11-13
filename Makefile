.PHONY: all deps fmt vet test cover build clean

all: deps fmt vet staticcheck test build

deps:
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...

staticcheck:
	staticcheck ./...

test:
	go test -coverprofile cover.out ./...

cover:
	go tool cover -html=cover.out

build:
	go build -o 6502_emulator main.go

clean:
	rm -f 6502_emulator cover.out
