.PHONY: all clean build run gen

HOST=127.0.0.1
PORT=8080
NAME=bookserver-micro

all: build run

build:
	go build -o bin/$(NAME)

run:
	bin/$(NAME) $(HOST) $(PORT)

clean:
	rm -f books/* db/sqlite3/books.db bin/*

gen:
	./gen.sh
