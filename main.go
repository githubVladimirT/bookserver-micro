package main

import (
	"encoding/json"
	"fmt"
	"io"

	"os"

	"strconv"

	"database/sql"
	"net/http"

	"github.com/blockloop/scan/v2"
	uuid "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"

	httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/register"
	"go.unistack.org/micro/v3/server"
)

// DB Queries
const (
	create_table = `
	CREATE TABLE IF NOT EXISTS "books" (
		"id"		INTEGER,
		"filepath"	TEXT NOT NULL,
		"title"		TEXT NOT NULL,
		"author"	TEXT NOT NULL,
		"genre"		TEXT NOT NULL,
		"year"		TEXT NOT NULL,
		PRIMARY 	KEY("id" AUTOINCREMENT)
	);
	`

	insert_book            = `INSERT INTO books VALUES(NULL, ?, ?, ?, ?, ?);`
	selectid_insertNewBook = `SELECT "id" FROM books WHERE filepath=$1`
	selectdata             = `SELECT "title", "author", "genre", "year" FROM "books" WHERE filepath=$1`
	sort_query             = `SELECT * FROM books ORDER BY `
	get_all_books          = `SELECT title FROM books;`
)

type Book struct {
	Filepath string `json:"filepath"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Year     string `json:"year"`
}

type ReturnBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}

type Status struct {
	Description string `json:"error"`
	Code        int    `json:"code"`
}

type StatusUploadedBook struct {
	Description string `json:"bookid"`
	Code        int    `json:"code"`
}

type DB struct{ db *sql.DB }

func (db *DB) insertNewBook(book Book) (int, error) {
	var id int

	_, err := db.db.Exec(insert_book, book.Filepath, book.Title, book.Author, book.Genre, book.Year)
	if err != nil {
		return -1, err
	}

	err = db.db.QueryRow(selectid_insertNewBook, book.Filepath).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// TODO: следующие функции имеют схожий функционал
func returnJSONStatus(w http.ResponseWriter, status string, code int) {
	text_error := Status{Description: status, Code: code}
	marshalled, err := json.Marshal(text_error)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func returnJSONStatusUploadedBook(w http.ResponseWriter, status string, code int) {
	text_error := StatusUploadedBook{Description: status, Code: code}
	marshalled, err := json.Marshal(text_error)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func (db *DB) GET_NULL_NULL(w http.ResponseWriter) {
	var books []string

	rows, err := db.db.Query(get_all_books)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadRequest)
	}

	err = scan.Rows(&books, rows)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadRequest)
	}

	marshalled, err := json.Marshal(books)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func (db *DB) GET_NOTNULL_NULL(w http.ResponseWriter, book string) {
	var title, author, genre, year string

	err := db.db.QueryRow(selectdata, "books/"+book).Scan(&title, &author, &genre, &year)
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal()
	}

	mybook := ReturnBook{Title: title, Author: author, Genre: genre, Year: year}
	marshalled, err := json.Marshal(mybook)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func (db *DB) GET_NULL_NOTNULL(w http.ResponseWriter, sort_type string) {
	books := []ReturnBook{}

	rows, err := db.db.Query(sort_query + sort_type)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadRequest)
	}

	err = scan.Rows(&books, rows)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadRequest)
	}

	marshalled, err := json.Marshal(books)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func PUSH_method_get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./templates/push.html")
	} else {
		returnJSONStatus(w, "invalid method", http.StatusBadGateway)
	}
}

func (db *DB) PUSH(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("book")
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
		return
	}

	book_title := r.FormValue("book_title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	year := r.FormValue("year")

	defer file.Close()

	filename := uuid.New().String()

	log.Info().Msg("--| New file |--")
	log.Info().Msg("Uploaded File: " + filename)
	log.Info().Msg("File Size: " + strconv.Itoa(int(handler.Size)))

	bookFile, err := os.CreateTemp("books", filename+"-*.pdf")
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer bookFile.Close()

	_, err = io.Copy(bookFile, file)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
		return
	}

	book := Book{
		Filepath: bookFile.Name(),
		Title:    book_title,
		Author:   author,
		Genre:    genre,
		Year:     year,
	}

	book_id, err := db.insertNewBook(book)
	if err != nil {
		returnJSONStatus(w, err.Error(), http.StatusBadGateway)
		return
	}

	returnJSONStatusUploadedBook(w, strconv.Itoa(book_id), http.StatusCreated)
}

func (db *DB) GET(w http.ResponseWriter, r *http.Request) {
	book := r.URL.Query().Get("book")
	sort_type := r.URL.Query().Get("sort")

	if book == "" && sort_type == "" {
		// book: NULL; sort_type: NULL
		db.GET_NULL_NULL(w)
	} else if book != "" && sort_type == "" {
		// book: NOT NULL; sort_type: NULL
		db.GET_NOTNULL_NULL(w, book)
	} else if book == "" && sort_type != "" {
		// book: NULL; sort_type: NOT NULL
		db.GET_NULL_NOTNULL(w, sort_type)
	} else if book != "" && sort_type != "" {
		// book: NOT NULL; sort_type: NOT NULL
		returnJSONStatus(w, "invalid argument combination", http.StatusBadRequest)
	}
}

func HOME(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/home.html")
}

func get_address() string {
	args := os.Args[1:]
	return args[0] + ":" + args[1]
}

// <----| MAIN FUNC |---->
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	db, err := sql.Open("sqlite3", "./db/books.db")
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal()
	}

	defer db.Close()

	_, err = db.Exec(create_table)
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal()
	}

	fmt.Fprintln(os.Stderr, "<----|   started   |---->")

	DataBase := &DB{db: db}

	reg := register.NewRegister()

	srv := httpsrv.NewServer(
		server.Address(get_address()),
		server.Name("bookserver"),
		server.Register(reg),
		httpsrv.PathHandler(http.MethodGet, "/", HOME),
		httpsrv.PathHandler(http.MethodGet, "/books", DataBase.GET),

		httpsrv.PathHandler(http.MethodPost, "/push", DataBase.PUSH),
		httpsrv.PathHandler(http.MethodGet, "/push", PUSH_method_get),
	)

	svc := micro.NewService(micro.Server(srv))
	svc.Init()
	svc.Run()

	fmt.Fprintln(os.Stderr, "<----| interrupted |---->")
}
