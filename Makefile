name := auth-forge
version := 0.1.0

sync:
	go mod tidy

run:
	@start-docker
	go run ./cmd/main.go

start-docker:
	docker compose up -d

format:
	gofmt -s -w .

lint:
	# go vet ./...
	@command -v golangci-lint >/dev/null 2>&1 \
    || sudo snap install golangci-lint --classic
	golangci-lint run

build: 
	go build ./cmd/main.go

