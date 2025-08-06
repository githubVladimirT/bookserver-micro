.PHONY: all clean init gen build run

HOST=127.0.0.1
PORT=8080
NAME=bookserver-micro

all: clean init build run

clean:
	rm -f books/* db/sqlite3/books.db bin/*

init:
	mkdir -p db/sqlite3/ books/ bin/

gen:
	./gen.sh

build:
	go build -o bin/$(NAME)

run:
	bin/$(NAME) $(HOST) $(PORT)
