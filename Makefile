.PHONY: build
BIN=out/copyhelper
MAIN=main.go
build:
	@echo "Building..."
	@go build -o $(BIN) $(MAIN)

windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BIN).exe $(MAIN)
