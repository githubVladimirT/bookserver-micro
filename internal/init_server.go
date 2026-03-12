package internal

import (
	"context"
	"time"

	"os"
	"path/filepath"
	"runtime"

	"github.com/githubVladimirT/bookserver-micro/http/handler"
	// pb "github.com/githubVladimirT/bookserver-micro/proto"

	// mhttp "go.unistack.org/micro-client-http/v3"
	// httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v4"
	"go.unistack.org/micro/v4/server"

	micro_logger "go.unistack.org/micro/v4/logger/slog"
)

func InitServerWithReady(ctx context.Context, readyCh chan<- struct{}) micro.Service {
	logger := micro_logger.NewLogger()
	if err := logger.Init(); err != nil {
		panic(err)
	}

	err := InitDirectories()
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	srv := server.NewServer(
		server.Name("bookserver-micro"),
		server.Context(ctx),
	)

	hd := srv.NewHandler(handler.NewServerHandler())

	if err := srv.Handle(hd); err != nil {
		logger.Fatal(ctx, err.Error())
	}

	service := micro.NewService(
		micro.Server(srv),
	)

	if err := service.Init(); err != nil {
		logger.Fatal(ctx, err.Error())
	}

	go func() {
		if err := service.Run(); err != nil {
			logger.Fatal(ctx, err.Error())
		}
	}()

	for !service.Ready() {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Millisecond):
		}
	}

	if readyCh != nil {
		close(readyCh)
	}

	return service
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
