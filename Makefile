name := auth-forge
version := 0.1.0

sync:
	go mod tidy

run:
	start-redis
	go run ./cmd/main.go

start-redis:
	docker compose up -d
