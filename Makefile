NAME="godu"
ARGS=""

build:
	go build -o bin/${NAME} *.go

run:
	./${NAME} ${ARGS}

clean:
	rm -rf ./bin/ *.out

test:
	go test

test_coverage:
	go test -coverprofile cover.out
	go tool cover -html=cover.out

lint:
	golangci-lint run --enable-all

all: lint test test_coverage build
