package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	pb "github.com/githubVladimirT/bookserver-micro/proto"

	"github.com/blockloop/scan/v2"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"

	// "github.com/rs/zerolog"
	"github.com/jmoiron/sqlx"
	log "github.com/rs/zerolog/log"
	httpsrv "go.unistack.org/micro-server-http/v3"
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

	insert_book            = `INSERT INTO books VALUES(NULL, "%s", "%s", "%s", "%s", "%s");`
	selectid_insertNewBook = `SELECT id FROM books WHERE filepath="%s";`
	selectdata             = `SELECT title, author, genre, year FROM books WHERE title="%s";`
	sort_query             = `SELECT title, author, genre, year FROM books ORDER BY %s;`
	get_all_books          = `SELECT title FROM books;`
)

type ReturnBook struct {
	Title  string `json:"filepath"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}

type Book struct {
	Filepath string `json:"filepath"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Year     string `json:"year"`
}

type ServerHandler struct {
	db *sqlx.DB
}

func InitDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "db/sqlite3/books.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(create_table)
	if err != nil {
		panic(err)
	}

	println("DB inited successfully")

	return db
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{db: InitDB()}
}

func (h *ServerHandler) InsertNewBook(book Book) (int, error) {
	var id int

	_, err := h.db.Exec(fmt.Sprintf(insert_book, book.Filepath, book.Title, book.Author, book.Genre, book.Year))
	if err != nil {
		return -1, err
	}

	err = h.db.QueryRow(fmt.Sprintf(selectid_insertNewBook, book.Filepath)).Scan(&id)
	if err != nil {
		return -1, err
	}

	println(id)

	return id, nil
}

func (h *ServerHandler) Home(ctx context.Context, req *pb.Empty, rsp *pb.StatusRsp) error {
	rsp.Description = "HomePage"

	httpsrv.SetRspCode(ctx, http.StatusOK)
	return nil
}

func (h *ServerHandler) GetAllBooks(ctx context.Context, req *pb.Empty, rsp *pb.GetAllBooks) error {
	var books []string

	rows, err := h.db.Query(get_all_books)
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	err = scan.Rows(&books, rows)
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	rsp.Books = books
	httpsrv.SetRspCode(ctx, http.StatusOK)

	return nil
}

func (h *ServerHandler) GetAllBooksAndSort(ctx context.Context, req *pb.SortType, rsp *pb.GetAllBooksAndSort) error {
	allowedSortFields := map[string]bool{
		"title":  true,
		"author": true,
		"genre":  true,
		"year":   true,
	}

	sortField := "title" // default value
	if req.SortType != "" && allowedSortFields[req.SortType] {
		sortField = req.SortType
	}

	query := fmt.Sprintf(`SELECT title, author, genre, year FROM books ORDER BY %s`, sortField)
	log.Info().Str("query", query).Msg("Executing query")

	var books []*pb.GetBookRsp
	err := h.db.Select(&books, query)

	if err != nil {
		log.Error().Err(err).Msg("Database query failed")
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	if len(books) == 0 {
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "No books found"})
	}

	rsp.Books = books
	httpsrv.SetRspCode(ctx, http.StatusOK)

	return nil
}
func (h *ServerHandler) Book(ctx context.Context, req *pb.GetBook, rsp *pb.GetBookRsp) error {
	var books []pb.GetBookRsp

	rows, err := h.db.Queryx(fmt.Sprintf(selectdata, req.BookTitle))
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}
	defer rows.Close()

	for rows.Next() {
		var b pb.GetBookRsp
		if err := rows.StructScan(&b); err != nil {
			log.Info().Msg(err.Error())
			httpsrv.SetRspCode(ctx, http.StatusBadRequest)
			return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
		}
		books = append(books, b)
	}

	if len(books) == 0 {
		httpsrv.SetRspCode(ctx, http.StatusNotFound)
		return httpsrv.SetError(&pb.StatusRsp{Description: "book not found"})
	}

	rsp.Title = books[0].Title
	rsp.Author = books[0].Author
	rsp.Genre = books[0].Genre
	rsp.Year = books[0].Year

	httpsrv.SetRspCode(ctx, http.StatusOK)
	return nil
}

func (h *ServerHandler) Push(ctx context.Context, req *pb.PostBook, rsp *pb.StatusUploadedBookRsp) error {
	file := req.BookBytes

	book_title := req.BookTitle
	author := req.Author
	genre := req.Genre
	year := req.Year

	filename := uuid.New().String()

	log.Info().Msg("--| New file |--")
	log.Info().Msg("Uploaded File: " + filename)

	bookFile, err := os.CreateTemp("books", filename+"-*.pdf")
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadGateway)
		return httpsrv.SetError(&pb.StatusRsp{Description: "temp file creation error"})
	}
	defer bookFile.Close()

	_, err = io.Copy(bookFile, bytes.NewReader(file))
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadGateway)
		return httpsrv.SetError(&pb.StatusRsp{Description: "file creation error"})
	}

	book := Book{
		Filepath: bookFile.Name(),
		Title:    book_title,
		Author:   author,
		Genre:    genre,
		Year:     year,
	}

	book_id := -1

	book_id, err = h.InsertNewBook(book)
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	rsp.BookId = strconv.Itoa(book_id)

	httpsrv.SetRspCode(ctx, http.StatusCreated)
	return nil
}
