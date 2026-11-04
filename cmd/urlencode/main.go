package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/argium/urlencode-urldecode/internal/cli"
)

func main() {
	err := cli.Run(os.Args[1:], os.Stdin, os.Stdout, func(input string) (string, error) {
		return url.QueryEscape(input), nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "urlencode: %v\n", err)
		os.Exit(1)
	}
}
