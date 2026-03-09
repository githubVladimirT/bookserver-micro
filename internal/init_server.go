package internal

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/githubVladimirT/bookserver-micro/handler"
	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v3"

	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/logger"
	"go.unistack.org/micro/v3/server"

	jsonpbcodec "go.unistack.org/micro-codec-jsonpb/v3"
)

func InitServer() {
	InitServerWithReady(nil)
}

func InitServerWithReady(readyCh chan<- struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := InitDirectories()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	options := append([]micro.Option{},
		micro.Server(httpsrv.NewServer(
			server.Name("bookserver-micro"),
			server.Version("1.0"),
			server.Address(":8080"),
			server.Context(ctx),
			server.Codec("application/json", jsonpbcodec.NewCodec()),
		)),

		micro.Client(mhttp.NewClient(
			client.Name("bookserver-micro-client"),
			client.Context(ctx),
			client.Codec("application/json", jsonpbcodec.NewCodec()),
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

	if readyCh != nil {
		close(readyCh)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}

func InitDirectories() error {
	projectRoot := GetProjectRoot()
	books_path := filepath.Join(projectRoot, "books/")
	db_path := filepath.Join(projectRoot, "db/sqlite3/")

	err := os.MkdirAll(books_path, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(db_path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func GetProjectRoot() string {
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
