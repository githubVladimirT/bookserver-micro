package main

import (
	"fmt"
	"os"

	"github.com/githubVladimirT/bookserver-micro/internal"
)

func main() {
	fmt.Fprintln(os.Stderr, "<----|   started   |---->")

	internal.InitServer()

	fmt.Fprintln(os.Stderr, "<----| interrupted |---->")
}
