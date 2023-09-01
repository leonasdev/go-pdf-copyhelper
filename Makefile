.PHONY: build
BIN=out/copyhelper
MAIN=main.go
build:
	@echo "Building..."
	@go build -o $(BIN) $(MAIN)
