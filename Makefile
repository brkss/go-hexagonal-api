
NAME=dextrace


all:
	go build cmd/main.go -o bin/$(NAME)


create-migration:
	migrate create -ext sql -dir internal/adapter/storage/postgres/migrations

