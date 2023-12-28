package handler

import (
	// "context"
	"testing"
)

var SERVER_HANDLER ServerHandler

func TestInitDB(t *testing.T) {
	db := InitDB()
	if db == nil {
		t.Errorf("InitDB() failed: %v", db)
	}
	SERVER_HANDLER = ServerHandler{db: db}
}

// func TestHome(t *testing.T) {
// 	req := &Empty{}
// 	rsp := &StatusRsp{}

// 	err := SERVER_HANDLER.Home(context.Background(), req, rsp)
// 	if err != nil {
// 		t.Errorf("Home() failed: %v", err)
// 	}

// 	if rsp.Description != "this is the home page" {
// 		t.Errorf("Home() incorrect response: %v", rsp)
// 	}
// }

// func TestGetAllBooks(t *testing.T) {
// 	req := &Empty{}
// 	rsp := &GetAllBooks{}

// 	if err := SERVER_HANDLER.GetAllBooks(context.Background(), req, rsp); err != nil {
// 		t.Errorf("GetAllBooks() failed: %v", err)
// 	}

// 	if len(rsp.Books) == 0 {
// 		t.Error("GetAllBooks() returned no books")
// 	}
// }

// func TestGetAllBooksAndSort(t *testing.T) {
// 	req := &SortType{SortType: "by_title"}
// 	rsp := &GetAllBooksAndSort{}

// 	if err := SERVER_HANDLER.GetAllBooksAndSort(context.Background(), req, rsp); err != nil {
// 		t.Errorf("GetAllBooksAndSort() failed: %v", err)
// 	}

// 	if len(rsp.Books) == 0 {
// 		t.Error("GetAllBooksAndSort() returned no books")
// 	}
// }

// func TestPush(t *testing.T) {
// 	req := &PostBook{
// 		file:       bytes("test"),
// 		book_title: "Test Book",
// 		author:     "Test Author",
// 		genre:      "Test Genre",
// 		year:       2023,
// 	}

// 	rsp := &StatusUploadedBookRsp{}
// 	if err := SERVER_HANDLER.Push(context.Background(), req, rsp); err != nil {
// 		t.Errorf("Push() failed: %v", err)
// 	}
// }

// func TestBook(t *testing.T) {
// 	req := &GetBook{}
// 	rsp := &GetBookRsp{}
// 	ans := &GetBookRsp{book_title: "Test Book",
// 		author: "Test Author",
// 		genre:  "Test Genre",
// 		year:   2023,
// 	}

// 	if err := SERVER_HANDLER.Book(context.Background(), req, rsp); err != nil {
// 		t.Errorf("Book() failed: %v", err)
// 	}
// 	if rsp != ans {
// 		t.Errorf("Book() incorrect response: %v", rsp)
// 	}
// }
