
.Phony: env lint build run proto

env:
ifeq ($(OS),Windows_NT)
	@echo "Setting up environment for Windows"
	@powershell.exe -File dotenv.ps1
else
	@echo "Setting up environment for Linux/Mac"
	@bash dotenv.sh
endif
	@echo "Environment setup finished"

lint:
	@echo "Running linter"
	@golangci-lint run cmd/... internal/... --config=./.golangci.yml
	@echo "Linter finished"

test:
	@echo "Running tests"
	go test ./internal/api/url -v
	@echo "Tests finished"

bench:
	@echo "Running benbenchmarks"
	go test ./internal/api/url -bench=. -benchmem
	@echo "Benchmarks finished"

build:
	@echo "Building application"
	@go build -o bin/ cmd/grpc_server/main.go
	@echo "Application built"

start:
	@echo "Starting application"
	@./bin/main
	@echo "Application started"

run:
	@echo "Running application"
	@go run cmd/grpc_server/main.go
	@echo "Application finished"

proto:
	mkdir -p pkg/url_v1 && \
    protoc --proto_path=api/url_v1 \
           --go_out=pkg/url_v1 --go_opt=paths=source_relative \
           --go-grpc_out=pkg/url_v1 --go-grpc_opt=paths=source_relative \
           api/url_v1/url.proto

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    	
.PHONY: logs compose-up compose-down storage-up storage-down

logs:
	@echo "Running logs"
	@docker-compose logs
	@echo "Logs finished"

compose:
	make compose-up

compose-up:
	@echo "Running docker-compose up"
	@docker-compose up -d
	@echo "Docker-compose up finished"

compose-down:
	@echo "Running docker-compose down"
	@docker-compose down
	@echo "Docker-compose down finished"

storage:
	make storage-up

storage-up:
	@echo "Running storage up"
	@docker-compose up -d postgres redis
	@echo "Storage up finished"

storage-down:
	@echo "Running storage down"
	@docker-compose down postgres redis
	@echo "Storage down finished"

