.PHONY: deps
deps:
	go get github.com/golang/lint/golint

.PHONY: build
build: 
	go build -o type tomoyukikobayashi

.PHONY: test
test: 
	go test -v -cover ./...

.PHONY: cover
cover: 
	go test -coverprofile=mainprof tomoyukikobayashi
	go tool cover -html=mainprof

.PHONY: lint
lint: deps
	go vet ./...
	golint ./...

.PHONY: fmt
fmt: deps
	go fmt *.go
