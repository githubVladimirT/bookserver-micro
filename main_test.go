package main

import (
	"context"
	"fmt"
	"os"
	"slices"

	"path/filepath"
	"testing"

	pb "github.com/githubVladimirT/bookserver-micro/http/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	jsoncodec "go.unistack.org/micro-codec-json/v3"
	"go.unistack.org/micro/v3/client"

	"github.com/githubVladimirT/bookserver-micro/internal"

	mregister "go.unistack.org/micro/v3/register/memory"
)

var (
	projectRoot  = internal.GetProjectRoot()
	TestBookPath = filepath.Join(projectRoot, "testbooks/TestBookBytes.pdf")
)

func TestHome(t *testing.T) {
	reg := mregister.NewRegister()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	srv := internal.InitServerWithReady(ctx, reg)
	t.Cleanup(func() { srv.Stop() })

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
		client.ContentType("application/json"),
		client.Codec("text/plain", jsoncodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Fatalf("Client init failed: %v", err)
	}

	cli := pb.NewBookServerClient(
		"bookclient-micro",
		c,
	)

	rsp, err := cli.Home(
		ctx,
		&pb.EmptyReq{},
		client.WithAddress(srv.Options().Address),
	)

	if err != nil {
		t.Fatalf("Home call failed: %v", err)
	}

	if rsp.Description != "HomePage" {
		t.Errorf("invalid rsp received: %#+v", rsp)
		return
	}
}

func TestPush(t *testing.T) {
	reg := mregister.NewRegister()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	srv := internal.InitServerWithReady(ctx, reg)
	t.Cleanup(func() { srv.Stop() })

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
		client.ContentType("application/json"),
		client.Codec("text/plain", jsoncodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Fatalf("Client init failed: %v", err)
	}

	cli := pb.NewBookServerClient(
		"bookclient-micro",
		c,
	)

	book_bytes, err := os.ReadFile(TestBookPath)
	if err != nil {
		t.Fatalf("failed to read test book: %v", err)
	}

	books_req := []*pb.PostBookReq{
		{
			BookTitle: "TestBook",
			Author:    "TestAuthor",
			Genre:     "TestGenre",
			Year:      "2023",
			BookBytes: book_bytes,
		},
		{
			BookTitle: "TestBook1",
			Author:    "TestAuthor1",
			Genre:     "TestGenre1",
			Year:      "2022",
			BookBytes: book_bytes,
		},
		{
			BookTitle: "TestBook2",
			Author:    "TestAuthor2",
			Genre:     "TestGenre2",
			Year:      "2021",
			BookBytes: book_bytes,
		},
	}

	for i := range books_req {
		rsp, err := cli.Push(
			ctx,
			books_req[i],
			client.WithAddress(srv.Options().Address),
			mhttp.Method("POST"),
		)

		if err != nil {
			t.Fatalf("Push call failed: %v", err)
		}

		if rsp.BookId == "-1" {
			t.Fatalf("invalid rsp received: %#+v", rsp)
		}
	}
}

func TestBook(t *testing.T) {
	reg := mregister.NewRegister()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	srv := internal.InitServerWithReady(ctx, reg)
	t.Cleanup(func() { srv.Stop() })

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
		client.ContentType("application/json"),
		client.Codec("text/plain", jsoncodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Fatalf("Client init failed: %v", err)
	}

	cli := pb.NewBookServerClient(
		"bookclient-micro",
		c,
	)

	rsp, err := cli.Book(
		ctx,
		&pb.GetBookReq{BookTitle: "TestBook"},
		client.WithAddress(srv.Options().Address),
	)

	if err != nil {
		t.Fatalf("Book call failed: %v", err)
	}

	if rsp.Title != "TestBook" || rsp.Author != "TestAuthor" || rsp.Genre != "TestGenre" || rsp.Year != "2023" {
		t.Fatalf("invalid rsp received: %#+v", rsp)
	}

	t.Logf("Book: %+v", rsp)
}

func TestGetAllBooks(t *testing.T) {
	reg := mregister.NewRegister()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	srv := internal.InitServerWithReady(ctx, reg)
	t.Cleanup(func() { srv.Stop() })

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
		client.ContentType("application/json"),
		client.Codec("text/plain", jsoncodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Fatalf("Client init failed: %v", err)
	}

	cli := pb.NewBookServerClient(
		"bookclient-micro",
		c,
	)

	rsp, err := cli.GetAllBooks(
		ctx,
		&pb.EmptyReq{},
		client.WithAddress(srv.Options().Address),
	)

	if err != nil {
		t.Fatalf("GetAllBooks call failed: %v", err)
	}

	if !slices.Contains(rsp.Books, "TestBook") {
		t.Errorf("invalid rsp received: %#+v", rsp)
		return
	}

	for i, b := range rsp.Books {
		t.Logf("Book %d: %+v", i, b)
	}
}

func TestGetAllBooksAndSort(t *testing.T) {
	reg := mregister.NewRegister()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	srv := internal.InitServerWithReady(ctx, reg)
	t.Cleanup(func() { srv.Stop() })

	c := mhttp.NewClient(
		client.Codec("application/json", jsoncodec.NewCodec()),
		client.ContentType("application/json"),
		client.Codec("text/plain", jsoncodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Fatalf("Client init failed: %v", err)
	}

	cli := pb.NewBookServerClient(
		"bookclient-micro",
		c,
	)

	book, err := cli.Book(
		ctx,
		&pb.GetBookReq{BookTitle: "TestBook"},
		client.WithAddress(srv.Options().Address),
	)

	if err != nil {
		t.Fatalf("GetAllBooksAndSort call failed: %v", err)
	}

	sortTypes := []string{"title", "author", "genre", "year"}
	for _, sortType := range sortTypes {
		rsp, err := cli.GetAllBooksAndSort(
			ctx,
			&pb.SortTypeReq{SortType: sortType},
			client.WithAddress(srv.Options().Address),
		)

		if err != nil {
			t.Fatalf("GetAllBooksAndSort call failed: %v", err)
		}

		if !containsBook(rsp.Books, book) {
			t.Errorf("book not found in %s sorted response", sortType)
			t.Logf("Expected book: %+v", book)
		}

		for i, b := range rsp.Books {
			t.Logf("Book %d: %+v", i, b)
		}
		fmt.Println("")
	}
}

func containsBook(books []*pb.GetBookRsp, target *pb.GetBookRsp) bool {
	for _, b := range books {
		if b.Title == target.Title &&
			b.Author == target.Author &&
			b.Genre == target.Genre &&
			b.Year == target.Year {
			return true
		}
	}
	return false
}
