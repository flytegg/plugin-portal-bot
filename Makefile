include .env

linux:
	@echo Building to linux...
	go env -w GOOS=linux
	go env -w GOARCH=amd64
	go build -o ./bin/bot/${LINUX_BINARY} ./cmd/main.go
	go env -u GOOS
	go env -u GOARCH

build:
	@echo Building...
	go build -o ./bin/bot/${BINARY} ./cmd/main.go

start:
	./bin/bot/${BINARY}

deploy: linux
	@echo Deploying...
	carbon deploy

restart: build start