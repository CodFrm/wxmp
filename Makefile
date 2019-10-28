
all: test cover build

build:
    go build .

test:
    go test -v ./...

cover:
	go test -cover -coverprofile=coverage.data ./...

cover-html: cover
	go tool cover -html=coverage.data -o coverage.html

