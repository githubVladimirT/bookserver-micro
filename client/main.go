package main

import (
	"context"
	"os"

	// "github.com/searKing/golang/go/exp/slices"

	"slices"

	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	jsonpbcodec "go.unistack.org/micro-codec-jsonpb/v3"
	"go.unistack.org/micro/v3/client"
	log "go.unistack.org/micro/v3/logger"
)

func main() {
	println("START HomeTest()")
	HomeTest()
	println("END HomeTest()")

	println("START PushTest()")
	PushTest()
	println("END PushTest()")

	println("START GetAllBooksTest()")
	GetAllBooksTest()
	println("END GetAllBooksTest()")

	println("START BookTest()")
	BookTest()
	println("END BookTest()")

	println("START GetAllBooksAndSortTest()")
	GetAllBooksAndSortTest()
	println("END GetAllBooksAndSortTest()")
}

func HomeTest() {
	logger := log.NewLogger(log.WithCallerSkipCount(2))
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		logger.Fatal(context.TODO(), err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress("http://127.0.0.1:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	req := &pb.Empty{}
	rsp, err := mc.Home(context.TODO(), req)
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	if rsp.Description != "HomePage" {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rsp)
		return
	}

	logger.Info(context.TODO(), "--| TEST PASSED |--")
}

func PushTest() {
	logger := log.NewLogger(log.WithCallerSkipCount(2))
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		logger.Fatal(context.TODO(), err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress("http://127.0.0.1:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	book_bytes, err := os.ReadFile("../TestBookBytes.pdf")
	if err != nil {
		logger.Error(context.TODO(), err)
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
		logger.Error(context.TODO(), err)
		return
	}

	if rspp.BookId == "-1" {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspp)
		return
	}

	logger.Info(context.TODO(), "--| TEST PASSED |--")
}

func BookTest() {
	logger := log.NewLogger(log.WithCallerSkipCount(2))
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		logger.Fatal(context.TODO(), err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress("http://127.0.0.1:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	rspb, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	if rspb.Title != "TestBook" || rspb.Author != "TestAuthor" || rspb.Genre != "TestGenre" || rspb.Year != "2023" {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspb)
		return
	}

	logger.Info(context.TODO(), "--| TEST PASSED |--")
}

func GetAllBooksTest() {
	logger := log.NewLogger(log.WithCallerSkipCount(2))
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		logger.Fatal(context.TODO(), err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress("http://127.0.0.1:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	rspgab, err := mc.GetAllBooks(context.TODO(), &pb.Empty{})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	if !slices.Contains(rspgab.Books, "TestBook") {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspgab)
		return
	}

	logger.Info(context.TODO(), "--| TEST PASSED |--")
}

func GetAllBooksAndSortTest() {
	logger := log.NewLogger(log.WithCallerSkipCount(2))
	c := mhttp.NewClient(
		client.ContentType("application/json"),
		client.Codec("application/json", jsonpbcodec.NewCodec()),
	)

	if err := c.Init(); err != nil {
		logger.Fatal(context.TODO(), err)
	}

	cli := client.NewClientCallOptions(
		c,
		client.WithAddress("http://127.0.0.1:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	rspgabas1, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "title"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}
	rspgabas2, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "author"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}
	rspgabas3, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "genre"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}
	rspgabas4, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "year"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	book, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	if !slices.Contains(rspgabas1.Books, book) {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspgabas1)
		return
	}
	if !slices.Contains(rspgabas2.Books, book) {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspgabas2)
		return
	}
	if !slices.Contains(rspgabas3.Books, book) {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspgabas3)
		return
	}
	if !slices.Contains(rspgabas4.Books, book) {
		logger.Errorf(context.TODO(), "invalid rsp received: %#+v", rspgabas4)
		return
	}

	logger.Info(context.TODO(), "--| TEST PASSED |--")
}
