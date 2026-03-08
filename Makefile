.PHONY: all clean init gen build run

HOST=127.0.0.1
PORT=8080
NAME=bookserver-micro
GO_CMD := go
GO_TEST := $(GO_CMD) test
GO_BUILD := $(GO_CMD) build

all: clean init build run

clean:
	rm -f books/* db/sqlite3/books.db bin/* coverage.out

init:
	mkdir -p db/sqlite3/ books/ bin/

gen:
	./gen.sh

build: init gen
	$(GO_BUILD) -o bin/$(NAME)

run: build
	bin/$(NAME) $(HOST) $(PORT)

test:
	$(GO_TEST) -v -race ./...

test-coverage:
	$(GO_TEST) -coverprofile=coverage.out ./...
	$(GO_CMD) tool cover -html=coverage.out
