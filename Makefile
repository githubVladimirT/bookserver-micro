.PHONY: all clean build run gen

HOST=127.0.0.1
PORT=8080

all: build run

build:
	go build .

run:
	./bookserver-micro $(HOST) $(PORT)

clean:
	rm -f books/* db/books.db

gen:
	./gen.sh
