GOFILES = $(shell find . -name '*.go')
GOPACKAGES = $(shell go list ./...)

test: $(GOFILES)
	go test -v -race ./...
