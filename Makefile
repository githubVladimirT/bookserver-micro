.PHONY: all clean init gen build run

HOST=127.0.0.1
PORT=8080
NAME=bookserver-micro
GO_CMD := go
GO_TEST := $(GO_CMD) test
GO_BUILD := $(GO_CMD) build
BOOKS_DIR := ./books/
DB_DIR := ./db/sqlite3/
BIN_DIR := ./bin/


all: clean init build run

clean:
	rm -f $(BOOKS_DIR)* $(DB_DIR)books.db $(BIN_DIR)* coverage.out

init:
	mkdir -p $(DB_DIR) $(BOOKS_DIR) $(BIN_DIR)

gen:
	cd ./http/ && go generate

build: init gen
	$(GO_BUILD) -o bin/$(NAME)

run: build
	bin/$(NAME) $(HOST) $(PORT)

test:
	$(GO_TEST) -v -race ./...

test-coverage:
	$(GO_TEST) -coverprofile=coverage.out ./...
	$(GO_CMD) tool cover -html=coverage.out
