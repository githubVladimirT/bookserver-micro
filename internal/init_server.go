package internal

import (
	"context"
	// "fmt"
	// "net"
	"os"
	"path/filepath"
	"runtime"

	// "strconv"
	// "strings"

	"github.com/githubVladimirT/bookserver-micro/handler"
	pb "github.com/githubVladimirT/bookserver-micro/proto"

	mhttp "go.unistack.org/micro-client-http/v3"
	httpsrv "go.unistack.org/micro-server-http/v3"
	"go.unistack.org/micro/v3"

	"go.unistack.org/micro/v3/client"
	micro_logger "go.unistack.org/micro/v3/logger"
	"go.unistack.org/micro/v3/server"

	jsonpbcodec "go.unistack.org/micro-codec-jsonpb/v3"
)

func InitServer() {
	InitServerWithReady(nil)
}

func InitServerWithReady(readyCh chan<- struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := micro_logger.NewLogger()

	err := InitDirectories()
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	// address := findAvailablePort(":8080")

	options := append([]micro.Option{},
		micro.Server(httpsrv.NewServer(
			server.Name("bookserver-micro"),
			server.Version("1.0"),
			server.Address(":8080"), // address
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
		logger.Fatal(ctx, err.Error())
	}

	eh := handler.NewServerHandler()

	if err := pb.RegisterBookServerServer(srv.Server(), eh); err != nil {
		logger.Fatal(ctx, err.Error())
	}

	go func() {
		if err := srv.Run(); err != nil {
			logger.Fatal(ctx, err.Error())
		}
	}()

	if readyCh != nil {
		close(readyCh)
	}

	<-ctx.Done()
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

// func findAvailablePort(initialAddress string) string {
// 	host, portStr, err := net.SplitHostPort(initialAddress)
// 	if err != nil {
// 		return initialAddress
// 	}

// 	port, err := strconv.Atoi(portStr)
// 	if err != nil {
// 		return initialAddress
// 	}

// 	for i := 0; i < 100; i++ {
// 		testAddress := fmt.Sprintf("%s:%d", host, port+i)
// 		ln, err := net.Listen("tcp", testAddress)
// 		if err == nil {
// 			ln.Close()
// 			return testAddress
// 		}

// 		if strings.Contains(err.Error(), "address already in use") ||
// 			strings.Contains(err.Error(), "bind: address already in use") {
// 			continue
// 		}

// 		return testAddress
// 	}

// 	return fmt.Sprintf("%s:%d", host, port)
// }
