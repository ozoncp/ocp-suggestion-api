PROJECT_NAME = ocp-suggestion-api

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/${PROJECT_NAME} ./cmd/${PROJECT_NAME}

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

.PHONY: generate
generate:
	mkdir -p swagger
	mkdir -p pkg/${PROJECT_NAME}
	protoc -I vendor.protogen \
		--go_out=pkg/${PROJECT_NAME} --go_opt=paths=import \
		--go-grpc_out=pkg/${PROJECT_NAME} --go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/${PROJECT_NAME} \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		--validate_out lang=go:pkg/${PROJECT_NAME} \
		--swagger_out=allow_merge=true,merge_file_name=api:swagger \
		api/${PROJECT_NAME}/${PROJECT_NAME}.proto
	mv pkg/${PROJECT_NAME}/github.com/ozoncp/${PROJECT_NAME}/pkg/${PROJECT_NAME}/* pkg/${PROJECT_NAME}/
	rm -rf pkg/${PROJECT_NAME}/github.com
	mkdir -p cmd/${PROJECT_NAME}
	cd pkg/${PROJECT_NAME} && ls go.mod || go mod init github.com/ozoncp/${PROJECT_NAME}/pkg/${PROJECT_NAME} && go mod tidy

.PHONY: vendor-proto
vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/${PROJECT_NAME}
	cp api/${PROJECT_NAME}/${PROJECT_NAME}.proto vendor.protogen/api/${PROJECT_NAME}
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/github.com/envoyproxy &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
	fi

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps:
	ls go.mod || go mod init github.com/ozoncp/${PROJECT_NAME}

.PHONY: deploy
deploy:
	docker-compose up -d

.PHONY: all
all: deps build
