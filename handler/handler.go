package handler

import (
	"bytes"
	"context"
	"database/sql"
	"io"
	"net/http"
	"os"
	"strconv"

	pb "github.com/githubVladimirT/bookserver-micro/proto"

	"github.com/blockloop/scan/v2"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	// "github.com/rs/zerolog"
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

	insert_book            = `INSERT INTO books VALUES(NULL, ?, ?, ?, ?, ?);`
	selectid_insertNewBook = `SELECT "id" FROM books WHERE filepath=$1`
	selectdata             = `SELECT "author", "genre", "year" FROM "books" WHERE title=?`
	sort_query             = `SELECT "id", "filepath", "title", "author", "genre", "year" FROM books ORDER BY ?`
	get_all_books          = `SELECT title FROM books;`
)

type ReturnBook struct {
	Title  string
	Author string
	Genre  string
	Year   string
}

type Book struct {
	Filepath string `json:"filepath"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Year     string `json:"year"`
}

type ServerHandler struct {
	db *sql.DB
}

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "db/sqlite3/books.db")
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

	_, err := h.db.Exec(insert_book, book.Filepath, book.Title, book.Author, book.Genre, book.Year)
	if err != nil {
		return -1, err
	}

	err = h.db.QueryRow(selectid_insertNewBook, book.Filepath).Scan(&id)
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
	var books []*pb.GetBookRsp

	print(sort_query + req.SortType)

	rows, err := h.db.Query(sort_query, req.SortType)
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

func (h *ServerHandler) Book(ctx context.Context, req *pb.GetBook, rsp *pb.GetBookRsp) error {
	var title, author, genre, year string

	title = req.BookTitle

	err := h.db.QueryRow(selectdata, title).Scan(&author, &genre, &year)
	if err != nil {
		log.Info().Msg(err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	rsp.Title = title
	rsp.Author = author
	rsp.Genre = genre
	rsp.Year = year

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
