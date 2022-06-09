//go:build ignore

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/adrianriobo/dater/pkg/schemas"
)

// os.Args[1] input filename
// os.Args[2] outout filename
// os.Args[3] package name
func main() {
	if err := schemas.GenerateFromJSONSchema(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
