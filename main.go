package main

import (
	"context"
	"fmt"
	"os"

	"github.com/githubVladimirT/bookserver-micro/handler"
	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v3"

	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/server"

	jsoncodec "go.unistack.org/micro-codec-json/v3"
	"go.unistack.org/micro/v3/logger"
)

func main() {
	fmt.Fprintln(os.Stderr, "<----|   started   |---->")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	options := append([]micro.Option{},
		micro.Server(httpsrv.NewServer(
			server.Name("bookserver-micro"),
			server.Version("1.0"),
			server.Address(":8080"),
			server.Context(ctx),
			server.Codec("application/json", jsoncodec.NewCodec()),
		)),

		micro.Client(mhttp.NewClient(
			client.Name("bookserver-micro-client"),
			client.Context(ctx),
			client.Codec("application/json", jsoncodec.NewCodec()),
			client.ContentType("application/json"),
		)),

		micro.Context(ctx),
	)

	srv := micro.NewService(options...)

	if err := srv.Init(); err != nil {
		logger.Fatal(ctx, err)
	}

	eh := handler.NewServerHandler()

	if err := pb.RegisterBookServerServer(srv.Server(), eh); err != nil {
		logger.Fatal(ctx, err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(ctx, err)
	}

	fmt.Fprintln(os.Stderr, "<----| interrupted |---->")
}
