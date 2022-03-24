PROGRAM_NAME = posop
BIN_DIR = bin


build:
	@echo "building $(BIN_DIR)/$(PROGRAM_NAME) executable"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(PROGRAM_NAME) main.go

vendor:
	@go mod vendor

tidy:
	@go mod tidy

run:
	@go run main.go

install:
	@echo "installing $(PROGRAM_NAME) executable"
	@go install

test:
	@go test -v ./...