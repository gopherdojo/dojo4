GOFILES = $(shell find . -name '*.go')
GOPACKAGES = $(shell go list ./...)

default: build

build: $(GOFILES)
	go build -o imgconv

test: $(GOFILES)
	go test -v -race ./...

cov: $(GOFILES)
	go test -coverprofile=cover.out -race ./...
	go tool cover -html=cover.out -o cover.html

lint: $(GOFILES)
	golangci-lint run -p bugs,unused,format,style ./...
