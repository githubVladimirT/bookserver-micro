package handler

import (
	"context"
	"os"
	"testing"

	"slices"

	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	jsonpbcodec "go.unistack.org/micro-codec-jsonpb/v3"
	"go.unistack.org/micro/v3/client"
)

const URL = "http://127.0.0.1:8080"
const TestBookPath = "./../testbooks/TestBookBytes.pdf"

func TestHome(t *testing.T) {
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Errorf(err.Error())
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress(URL),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	req := &pb.Empty{}
	rsp, err := mc.Home(context.TODO(), req)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if rsp.Description != "HomePage" {
		t.Errorf("invalid rsp received: %#+v", rsp)
		return
	}
}

func TestPush(t *testing.T) {
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Errorf(err.Error())
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress(URL),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	book_bytes, err := os.ReadFile(TestBookPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	book := &pb.PostBook{
		BookTitle: "TestBook",
		Author:    "TestAuthor",
		Genre:     "TestGenre",
		Year:      "2023",
		BookBytes: book_bytes,
	}

	rspp, err := mc.Push(context.TODO(), book)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if rspp.BookId == "-1" {
		t.Errorf("invalid rsp received: %#+v", rspp)
		return
	}
}

func TestBook(t *testing.T) {
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Errorf(err.Error())
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress(URL),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	rspb, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if rspb.Title != "TestBook" || rspb.Author != "TestAuthor" || rspb.Genre != "TestGenre" || rspb.Year != "2023" {
		t.Errorf("invalid rsp received: %#+v", rspb)
		return
	}
}

func TestGetAllBooks(t *testing.T) {
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		t.Errorf(err.Error())
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress(URL),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	rspgab, err := mc.GetAllBooks(context.TODO(), &pb.Empty{})
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !slices.Contains(rspgab.Books, "TestBook") {
		t.Errorf("invalid rsp received: %#+v", rspgab)
		return
	}
}

func TestGetAllBooksAndSort(t *testing.T) {
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)
	if err := c.Init(); err != nil {
		t.Fatalf("client init failed: %v", err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress(URL),
	)
	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	book, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
	if err != nil {
		t.Fatalf("failed to get test book: %v", err)
	}

	sortTypes := []string{"title", "author", "genre", "year"}
	for _, sortType := range sortTypes {
		rsp, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: sortType})
		if err != nil {
			t.Errorf("sort by %s failed: %v", sortType, err)
			continue
		}

		if !containsBook(rsp.Books, book) {
			t.Errorf("book not found in %s sorted response", sortType)
			t.Logf("Expected book: %+v", book)
			for i, b := range rsp.Books {
				t.Logf("Book %d: %+v", i, b)
			}
		}
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
