all: deps migrate run

migrate:
	go run app/tooling/admin/main.go migrate

deps:
	go mod tidy

test:
	go test ./... --cover

run:
	go run app/service/api/main.go

build:
	go build -o api app/service/api/main.go

