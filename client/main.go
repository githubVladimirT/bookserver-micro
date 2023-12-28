package main

import (
	"context"

	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	jsoncodec "go.unistack.org/micro-codec-json/v3"
	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/logger"
)

func main() {
	// cli := client.NewClientCallOptions(
	// 	mhttp.NewClient(
	// 		client.ContentType("application/json"),
	// 		client.Codec("application/json", jsoncodec.NewCodec()),
	// 	),
	// 	client.WithAddress("0.0.0.0:8080"),
	// )

	// mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	println("START HomeTest()")
	// HomeTest(mc)
	HomeTest()
	println("\nEND HomeTest()")

	// println("START PushTest()")
	// PushTest(mc)
	// println("END PushTest()")

	// println("START GetAllBooksTest()")
	// GetAllBooksTest(mc)
	// println("END GetAllBooksTest()")

	// println("START BookTest()")
	// BookTest(mc)
	// println("END BookTest()")

	// println("START GetAllBooksAndSortTest()")
	// GetAllBooksAndSortTest(mc)
	// println("END GetAllBooksAndSortTest()")
}

// func GetAllBooksAndSortTest(mc pb.BookServerClient) {
// 	rspgabas1, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "title"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}
// 	rspgabas2, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "author"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}
// 	rspgabas3, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "genre"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}
// 	rspgabas4, err := mc.GetAllBooksAndSort(context.TODO(), &pb.SortType{SortType: "year"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	book, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	if !slices.Contains(rspgabas1.Books, book) {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspgabas1)
// 		return
// 	}
// 	if !slices.Contains(rspgabas2.Books, book) {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspgabas2)
// 		return
// 	}
// 	if !slices.Contains(rspgabas3.Books, book) {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspgabas3)
// 		return
// 	}
// 	if !slices.Contains(rspgabas4.Books, book) {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspgabas4)
// 		return
// 	}
// }

// func HomeTest(mc pb.BookServerClient) {
// 	req := &pb.Empty{}
// 	rsph, err := mc.Home(context.TODO(), req)

// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	if rsph.Description != "HomePage" {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rsph)
// 		return
// 	}
// }

func HomeTest() {
	cli := client.NewClientCallOptions(
		mhttp.NewClient(
			client.ContentType("application/json"),
			client.Codec("application/json", jsoncodec.NewCodec()),
		),
		client.WithAddress("0.0.0.0:8080"),
	)

	mc := pb.NewBookServerClient("bookserver-micro-client", cli)

	req := &pb.Empty{}	
	rsp, err := mc.Home(context.TODO(), req)
	logger.Info(context.TODO(), cli.Options().ContentType)

	if err != nil {
		logger.Error(context.TODO(), err)
		return
	}

	if rsp.Description != "HomePage" {
		logger.Error(context.TODO(), "invalid rsp received: %#+v", rsp)
		return
	}
}

// func PushTest(mc pb.BookServerClient) {
// 	book_bytes, err := os.ReadFile("../books/TestBookBytes.pdf")
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	book := &pb.PostBook{
// 		BookTitle: "TestBook",
// 		Author:    "TestAuthor",
// 		Genre:     "TestGenre",
// 		Year:      "2023",
// 		BookBytes: book_bytes,
// 	}

// 	rspp, err := mc.Push(context.TODO(), book)
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	if rspp.BookId != "-1" {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspp)
// 		return
// 	}
// }

// func BookTest(mc pb.BookServerClient) {
// 	rspb, err := mc.Book(context.TODO(), &pb.GetBook{BookTitle: "TestBook"})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	if rspb.Title != "TestBook" || rspb.Author != "TestAuthor" || rspb.Genre != "TestGenre" || rspb.Year != "2023" {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspb)
// 		return
// 	}
// }

// func GetAllBooksTest(mc pb.BookServerClient) {
// 	rspgab, err := mc.GetAllBooks(context.TODO(), &pb.Empty{})
// 	if err != nil {
// 		logger.Error(context.TODO(), err)
// 		return
// 	}

// 	if !slices.Contains(rspgab.Books, "TestBook") {
// 		logger.Error(context.TODO(), "invalid rsp received: %#+v", rspgab)
// 		return
// 	}
// }
