package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "<----|   started   |---->")

	fmt.Fprintln(os.Stderr, "<----| interrupted |---->")
}
