.PHONY: migrate-up build-linux run

# migrate-up:
#     migrate -database "postgres://golang_migrate:P@ssw0rd@localhost:5432/db_cats?sslmode=disable" -path db/migrations up

build-linux:
    env GOARCH=amd64 GOOS=linux go build -o main app.go

run:
    go run app.go