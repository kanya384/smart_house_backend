include .env

export GOOSE_DRIVER := postgres
export GOOSE_DBSTRING := host=${APP_POSTGRES_HOST} port=${APP_POSTGRES_PORT} user=${APP_POSTGRES_USERNAME} password=${APP_POSTGRES_PASS} dbname=${APP_POSTGRES_DB_NAME} sslmode=disable

up:
	@docker-compose up -d & disown
run:
	@go run cmd/app/main.go
update:
	@go run cmd/command/main.go
migrate-up:
	@goose -dir=./internal/migrations up
migrate-down:
	@goose -dir=./internal/migrations down
proto-gen:
	@protoc --go_out=internal api/grpc/api.proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=internal