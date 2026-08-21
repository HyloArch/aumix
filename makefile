BINARY_NAME=aumix
BUILD_DIR=build

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./..

.PHONY: build-frontend
build-frontend:
	(cd frontend && npm run build)

.PHONY: run-frontend
run-frontend:
	(cd frontend && npm run dev)

.PHONY: build
build: tidy fmt
	mkdir -p $(BUILD_DIR)/$(BINARY_NAME)
	CC=aarch64-unknown-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)/$(BINARY_NAME)
	shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)/$(BINARY_NAME) > $(BUILD_DIR)/$(BINARY_NAME)/signature.txt

.PHONY: run
run: fmt build-frontend
	go run .

