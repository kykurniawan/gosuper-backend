run:
	@go run cmd/main.go

build:
	@go build -o bin/gosuper cmd/main.go

install-wire:
	@GOBIN=bin go install github.com/google/wire/cmd/wire

wire:
	@bin/wire gen gosuper/app