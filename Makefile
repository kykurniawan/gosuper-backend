build:
	@go build -o bin/gosuper cmd/main.go

run-dev:
	@go run cmd/main.go

run-prod:
	@go build -o bin/gosuper cmd/main.go
	@bin/gosuper

install-wire:
	@GOBIN=bin go install github.com/google/wire/cmd/wire

wire:
	@bin/wire gen gosuper/app