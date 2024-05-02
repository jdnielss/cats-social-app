.PHONY: migrate-up migrate-down build run

migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/db_cats?sslmode=disable" -path database/migration up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/db_cats?sslmode=disable" -path database/migration down

build:
	env GOARCH=amd64 GOOS=linux go build -o main app.go

run:
	go run app.go
