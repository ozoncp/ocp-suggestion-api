PROJECT_NAME = ocp-suggestion-api

.PHONY: build
build:
	go mod tidy
	go mod vendor
	go build -o bin/${PROJECT_NAME} ./cmd/${PROJECT_NAME}

.PHONY: run
run:
	go run ./cmd/${PROJECT_NAME}

.PHONY: test
test:
	go test -test.v -coverprofile=cover.out ./...

.PHONY: cover
cover:	test
	go tool cover -html=cover.out

.PHONY: benchmark
benchmark:
	go test -bench=. -benchtime=10000x -count=5 ./...

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: clean
clean:
	@rm -f cover.out
	@rm -rf bin

.PHONY: all
all: build
