package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	pb "github.com/githubVladimirT/bookserver-micro/http/proto"

	// "github.com/blockloop/scan/v2"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
	httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v3/logger"
	"go.unistack.org/micro/v3/logger/slog"
)

func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

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
	selectid_insertNewBook = `SELECT id FROM books WHERE filepath=?;`
	selectdata             = `SELECT title, author, genre, year FROM books WHERE title=?;`
	sort_query             = `SELECT title, author, genre, year FROM books ORDER BY ?;`
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

func InitLog() logger.Logger {
	log := slog.NewLogger()
	if err := log.Init(); err != nil {
		panic(err)
	}
	return log
}

func InitDB() *sqlx.DB {
	log := InitLog()
	projectRoot := getProjectRoot()
	dbPath := filepath.Join(projectRoot, "db/sqlite3/books.db")
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(context.Background(), err.Error())
	}

	_, err = db.Exec(create_table)
	if err != nil {
		log.Fatal(context.Background(), err.Error())
	}

	log.Info(context.Background(), "database initialized successfully")

	return db
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{db: InitDB()}
}

func (h *ServerHandler) Home(ctx context.Context, req *pb.EmptyReq, rsp *pb.StatusRsp) error {
	rsp.Description = "HomePage"

	httpsrv.SetRspCode(ctx, http.StatusOK)
	return nil
}

func (h *ServerHandler) GetAllBooks(ctx context.Context, req *pb.EmptyReq, rsp *pb.GetAllBooksRsp) error {
	log := InitLog()

	var books []string

	rows, err := h.db.Query(get_all_books)
	if err != nil {
		log.Error(ctx, err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Error(ctx, err.Error())
			httpsrv.SetRspCode(ctx, http.StatusBadRequest)
			return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
		}
		books = append(books, title)
	}

	if err := rows.Err(); err != nil {
		log.Error(ctx, err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	rsp.Books = books
	httpsrv.SetRspCode(ctx, http.StatusOK)

	return nil
}

func (h *ServerHandler) GetAllBooksAndSort(ctx context.Context, req *pb.SortTypeReq, rsp *pb.GetAllBooksAndSortRsp) error {
	log := InitLog()

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

	var books []*pb.GetBookRsp
	err := h.db.Select(&books, sort_query, sortField)

	if err != nil {
		log.Error(ctx, err.Error())
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
func (h *ServerHandler) Book(ctx context.Context, req *pb.GetBookReq, rsp *pb.GetBookRsp) error {
	log := InitLog()

	var title, author, genre, year string

	err := h.db.QueryRow(selectdata, req.BookTitle).Scan(&title, &author, &genre, &year)
	if err != nil {
		log.Error(ctx, err.Error())
		httpsrv.SetRspCode(ctx, http.StatusNotFound)
		return httpsrv.SetError(&pb.StatusRsp{Description: "book not found"})
	}

	rsp.Title = title
	rsp.Author = author
	rsp.Genre = genre
	rsp.Year = year

	httpsrv.SetRspCode(ctx, http.StatusOK)
	return nil
}

func (h *ServerHandler) Push(ctx context.Context, req *pb.PostBookReq, rsp *pb.StatusUploadedBookRsp) error {
	log := InitLog()

	file := req.BookBytes

	book_title := req.BookTitle
	author := req.Author
	genre := req.Genre
	year := req.Year

	filename := uuid.New().String()

	log.Info(ctx, "Uploaded File: "+filename)

	bookFile, err := os.CreateTemp("books", filename+"-*.pdf")
	if err != nil {
		log.Error(ctx, err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadGateway)
		return httpsrv.SetError(&pb.StatusRsp{Description: "temp file creation error"})
	}
	defer bookFile.Close()

	_, err = io.Copy(bookFile, bytes.NewReader(file))
	if err != nil {
		log.Error(ctx, err.Error())
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

	book_id, err = h.insertNewBook(book)
	if err != nil {
		log.Error(ctx, err.Error())
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
		return httpsrv.SetError(&pb.StatusRsp{Description: "database error"})
	}

	rsp.BookId = strconv.Itoa(book_id)

	httpsrv.SetRspCode(ctx, http.StatusCreated)
	return nil
}

func (h *ServerHandler) insertNewBook(book Book) (int, error) {
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
