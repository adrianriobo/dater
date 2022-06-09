//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/adrianriobo/dater/pkg/schemas"
)

// os.Args[1] input path
// os.Args[2] input filename
// os.Args[2] outout path
// os.Args[3] package name
func main() {
	if err := schemas.GenerateFromXSD(os.Args[1], os.Args[2], os.Args[3], os.Args[4]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
