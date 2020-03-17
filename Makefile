.PHONY: install test build clean

all: install

install:
	go install -v

test:
	go test -v ./...
	./go-cli print a b c

build: clean
	go vet ./...
	go fmt ./...
	go build

clean:
	-rm -rf ./go-cli
	go mod tidy
