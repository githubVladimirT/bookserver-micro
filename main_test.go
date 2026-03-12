package main

import (
	"context"
	// "fmt"
	// "os"
	"path/filepath"
	"testing"

	// "slices"

	pb "github.com/githubVladimirT/bookserver-micro/http/proto"

	mhttp "go.unistack.org/micro-client-http/v4"
	jsoncodec "go.unistack.org/micro-codec-json/v4"
	"go.unistack.org/micro/v4/client"

	"github.com/githubVladimirT/bookserver-micro/internal"
)

var (
	projectRoot  = internal.GetProjectRoot()
	TestBookPath = filepath.Join(projectRoot, "testbooks/TestBookBytes.pdf")
)

func TestHome(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readyCh := make(chan struct{})
	// service := internal.InitServerWithReady(ctx, readyCh)
	// if service == nil {
	// 	t.Fatal("Service start failed")
	// }
	// defer service.Stop()
	// <-readyCh

	service := internal.InitServerWithReady(ctx, readyCh)

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
	)

	if c == nil {
		t.Fatal("mhttp.NewClient returned nil")
	}

	req := c.NewRequest(
		"bookserver-micro",
		"/",
		&pb.EmptyReq{},
	)
	rsp := new(pb.StatusRsp)

	err := c.Call(
		ctx,
		req,
		rsp,
		client.WithAddress(service.Server().Options().Address),
		mhttp.Method("GET"),
	)

	if err != nil {
		t.Fatalf("Home call failed: %v", err)
	}

	if rsp.Description != "HomePage" {
		t.Errorf("invalid rsp received: %#+v", rsp)
		return
	}
}

// func TestPush(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	readyCh := make(chan struct{})
// 	server := internal.InitServerWithReady(ctx, readyCh)
// 	defer server.Stop()
// 	<-readyCh

// 	c := mhttp.NewClient(
// 		client.ContentType("application/json"),
// 		client.Codec("application/json",
// 			jsoncodec.NewCodec()),
// 	)

// 	if err := c.Init(); err != nil {
// 		t.Fatalf("client init failed: %v", err)
// 	}

// 	cli := client.NewClientCallOptions(
// 		c,
// 		client.WithAddress(server.Server().Options().Address),
// 	)

// 	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

// 	book_bytes, err := os.ReadFile(TestBookPath)
// 	if err != nil {
// 		t.Fatalf("failed to read test book: %v", err)
// 	}

// 	books := []*pb.PostBook{
// 		{
// 			BookTitle: "TestBook",
// 			Author:    "TestAuthor",
// 			Genre:     "TestGenre",
// 			Year:      "2023",
// 			BookBytes: book_bytes,
// 		},
// 		{
// 			BookTitle: "TestBook1",
// 			Author:    "TestAuthor1",
// 			Genre:     "TestGenre1",
// 			Year:      "2022",
// 			BookBytes: book_bytes,
// 		},
// 		{
// 			BookTitle: "TestBook2",
// 			Author:    "TestAuthor2",
// 			Genre:     "TestGenre2",
// 			Year:      "2021",
// 			BookBytes: book_bytes,
// 		},
// 	}

// 	for i := range books {
// 		book := books[i]

// 		rspp, err := mc.Push(context.TODO(), book)
// 		if err != nil {
// 			t.Fatalf("Push call failed: %v", err)
// 		}

// 		if rspp.BookId == "-1" {
// 			t.Errorf("invalid rsp received: %#+v", rspp)
// 			return
// 		}
// 	}
// }

// func TestBook(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	readyCh := make(chan struct{})
// 	server := internal.InitServerWithReady(ctx, readyCh)
// 	defer server.Stop()
// 	<-readyCh

// 	c := mhttp.NewClient(
// 		client.ContentType("application/json"),
// 		client.Codec("application/json", jsoncodec.NewCodec()),
// 	)

// 	if err := c.Init(); err != nil {
// 		t.Fatalf("client init failed: %v", err)
// 	}

// 	cli := client.NewClientCallOptions(
// 		c,
// 		client.WithAddress(server.Server().Options().Address),
// 	)

// 	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

// 	rspb, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
// 	if err != nil {
// 		t.Fatalf("Book call failed: %v", err)
// 	}

// 	if rspb.Title != "TestBook" || rspb.Author != "TestAuthor" || rspb.Genre != "TestGenre" || rspb.Year != "2023" {
// 		t.Errorf("invalid rsp received: %#+v", rspb)
// 		return
// 	}

// 	t.Logf("Book: %+v", rspb)
// }

// func TestGetAllBooks(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	readyCh := make(chan struct{})
// 	server := internal.InitServerWithReady(ctx, readyCh)
// 	defer server.Stop()
// 	<-readyCh

// 	c := mhttp.NewClient(
// 		client.ContentType("application/json"),
// 		client.Codec("application/json", jsoncodec.NewCodec()),
// 	)

// 	if err := c.Init(); err != nil {
// 		t.Fatalf("client init failed: %v", err)
// 	}

// 	cli := client.NewClientCallOptions(
// 		c,
// 		client.WithAddress(server.Server().Options().Address),
// 	)

// 	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

// 	rspgab, err := mc.GetAllBooks(context.TODO(), &pb.Empty{})
// 	if err != nil {
// 		t.Fatalf("GetAllBooks call failed: %v", err)
// 	}

// 	if !slices.Contains(rspgab.Books, "TestBook") {
// 		t.Errorf("invalid rsp received: %#+v", rspgab)
// 		return
// 	}

// 	for i, b := range rspgab.Books {
// 		t.Logf("Book %d: %+v", i, b)
// 	}
// }

// func TestGetAllBooksAndSort(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	readyCh := make(chan struct{})
// 	server := internal.InitServerWithReady(ctx, readyCh)
// 	defer server.Stop()
// 	<-readyCh

// 	c := mhttp.NewClient(
// 		client.ContentType("application/json"),
// 		client.Codec("application/json", jsoncodec.NewCodec()),
// 	)
// 	if err := c.Init(); err != nil {
// 		t.Fatalf("client init failed: %v", err)
// 	}

// 	cli := client.NewClientCallOptions(
// 		c,
// 		client.WithAddress(server.Server().Options().Address),
// 	)
// 	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

// 	book, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
// 	if err != nil {
// 		t.Fatalf("failed to get test book: %v", err)
// 	}

// 	sortTypes := []string{"title", "author", "genre", "year"}
// 	for _, sortType := range sortTypes {
// 		rsp, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: sortType})
// 		if err != nil {
// 			t.Errorf("sort by %s failed: %v", sortType, err)
// 			continue
// 		}

// 		if !containsBook(rsp.Books, book) {
// 			t.Errorf("book not found in %s sorted response", sortType)
// 			t.Logf("Expected book: %+v", book)
// 		}

// 		for i, b := range rsp.Books {
// 			t.Logf("Book %d: %+v", i, b)
// 		}
// 		fmt.Println("")
// 	}
// }

// func containsBook(books []*pb.GetBookRsp, target *pb.GetBookRsp) bool {
// 	for _, b := range books {
// 		if b.Title == target.Title &&
// 			b.Author == target.Author &&
// 			b.Genre == target.Genre &&
// 			b.Year == target.Year {
// 			return true
// 		}
// 	}
// 	return false
// }
