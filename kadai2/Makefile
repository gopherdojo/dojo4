.PHONY: deps
deps:
	go get github.com/golang/lint/golint

.PHONY: build
build: 
	go build -o imconv tomoyukikobayashi

.PHONY: test
test: 
	go test -v -cover ./...

.PHONY: cover
cover: 
	go test -coverprofile=main tomoyukikobayashi
	go tool cover -html=main
	go test -coverprofile=imconv tomoyukikobayashi/imconv
	go tool cover -html=imconv
	go test -coverprofile=file tomoyukikobayashi/file
	go tool cover -html=file

.PHONY: lint
lint: deps
	go vet ./...
	golint ./...

.PHONY: fmt
fmt: deps
	go fmt *.go
