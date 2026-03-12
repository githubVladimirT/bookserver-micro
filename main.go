package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/githubVladimirT/bookserver-micro/internal"
)

func main() {
	fmt.Fprintln(os.Stderr, "<----|   started   |---->")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readyCh := make(chan struct{})
	srv := internal.InitServerWithReady(ctx, readyCh)

	<-readyCh
	fmt.Fprintln(os.Stderr, "<----|   server ready   |---->")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Fprintln(os.Stderr, "<----| interrupted |---->")

	if srv != nil {
		srv.Stop()
	}
}
