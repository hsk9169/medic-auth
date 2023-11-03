.PHONY: init test mocks tidy docs dev-run
APP_NAME = auth

init: tidy docs 
	go mod tidy


test:
	go test ./... -v
	
tidy:
	go install github.com/vektra/mockery/v2@v2.32.4


mocks:
	mockery --all --keeptree

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/$(APP_NAME)/main.go -o docs

local-run:
	go run cmd/$(APP_NAME)/main.go --config=./configs/local/application.yaml

build: 
	go build -o $(APP_NAME) cmd/$(APP_NAME)/main.go

