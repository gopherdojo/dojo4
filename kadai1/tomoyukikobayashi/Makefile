.PHONY: deps
deps:
	go get github.com/golang/lint/golint

.PHONY: build
build: deps
	go build -o imconv

.PHONY: test
test: deps
	go test -v ./...

.PHONY: lint
lint: deps
	go vet ./...
	golint ./...

.PHONY: fmt
fmt: deps
	go fmt *.go