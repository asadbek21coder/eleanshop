CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

run:
	go run cmd/main.go

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:asadbek@localhost:5432/eleanshop?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:asadbek@localhost:5432/eleanshop?sslmode=disable' down
