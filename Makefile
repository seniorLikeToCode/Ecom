APP_NAME=Ecom

build: 
	@go build -o bin/$(APP_NAME) cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/$(APP_NAME)