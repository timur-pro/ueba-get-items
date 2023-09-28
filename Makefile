MAIN_FILE=./cmd/main.go
CMD_FILE=./cmd/get-item-service

build:
	go build -o $(CMD_FILE) $(MAIN_FILE)

test:
	go test -v -cover ./...

run: build
	$(CMD_FILE)

generate:
	mockgen -destination=./internal/handlers/mocks.go -source=./internal/handlers/repository.go -package=handlers