include .env
export APP_PORT
export APP_LOG_LEVEL
export APP_SALT
export APP_TOKEN_TTL
export APP_REFRESH_TOKEN_TTL
export APP_TOKEN_SECRET
export APP_POSTGRES_HOST
export APP_POSTGRES_PORT
export APP_POSTGRES_USERNAME
export APP_POSTGRES_PASS
export APP_POSTGRES_DB_NAME
export APP_POSTGRES_POOL_CONNECTIONS
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
mockgen:
	@cd internal/repository && mockery --all --recursive=true --output=../../mocks
run_tests:
	@goose -dir=./test/migrations up && go test ./... -v && goose -dir=./test/migrations down
