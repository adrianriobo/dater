//go:build ignore

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/adrianriobo/dater/pkg/schemas"
)

// os.Args[1] sourcePath
// os.Args[2] targetPath
// os.Args[3] package name
func main() {
	if err := schemas.StructFromJSONSchema(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
