package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/argium/urlencode-urldecode/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:], os.Stdin, os.Stdout, url.QueryUnescape); err != nil {
		fmt.Fprintf(os.Stderr, "urldecode: %v\n", err)
		os.Exit(1)
	}
}
