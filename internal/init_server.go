package internal

import (
	"context"
	"fmt"

	"os"
	"path/filepath"
	"runtime"

	"github.com/githubVladimirT/bookserver-micro/http/handler"
	pb "github.com/githubVladimirT/bookserver-micro/http/proto"
	httpsrv "go.unistack.org/micro-server-http/v3"

	"go.unistack.org/micro/v3/register"
	"go.unistack.org/micro/v3/server"

	jsoncodec "go.unistack.org/micro-codec-json/v3"

	micro_logger "go.unistack.org/micro/v3/logger/slog"
)

func InitServerWithReady(ctx context.Context, reg register.Register) *httpsrv.Server {
	logger := micro_logger.NewLogger()
	if err := logger.Init(); err != nil {
		panic(err)
	}

	err := InitDirectories()
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	srv := httpsrv.NewServer(
		server.Address("127.0.0.1:0"),
		server.Name("bookserver-micro"),
		server.Register(reg),
		server.Context(ctx),
		server.Codec("application/json", jsoncodec.NewCodec()),
	)

	if err := pb.RegisterBookServerServer(srv, handler.NewServerHandler()); err != nil {
		logger.Fatal(ctx, err.Error())
	}

	if err := srv.Start(); err != nil {
		logger.Fatal(ctx, err.Error())
	}

	svc, err := reg.LookupService(ctx, "bookserver-micro")
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	if len(svc) != 1 {
		logger.Fatal(ctx, fmt.Errorf("Expected 1 service got %d: %+v", len(svc), svc).Error())
	}

	if len(svc[0].Nodes) != 1 {
		logger.Fatal(ctx, fmt.Errorf("Expected 1 node got %d: %+v", len(svc[0].Nodes), svc[0].Nodes).Error())
	}

	return srv
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
