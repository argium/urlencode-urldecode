// Package cli provides the shared input and output handling for the URL commands.
package cli

import (
	"errors"
	"fmt"
	"io"
)

// Run transforms one argument, or all of stdin when no argument is supplied.
// It writes the result to stdout without adding a newline. Multiple arguments,
// read failures, transformation failures and write failures return an error.
// Run does not close stdin or stdout.
func Run(args []string, stdin io.Reader, stdout io.Writer, transform func(string) (string, error)) error {
	if len(args) > 1 {
		return errors.New("expected at most one argument; quote input containing spaces")
	}

	var input string
	if len(args) == 1 {
		input = args[0]
	} else {
		data, err := io.ReadAll(stdin)
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		input = string(data)
	}

	output, err := transform(input)
	if err != nil {
		return err
	}
	if _, err := io.WriteString(stdout, output); err != nil {
		return fmt.Errorf("write stdout: %w", err)
	}
	return nil
}
