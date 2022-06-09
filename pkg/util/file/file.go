package file

import (
	"fmt"
	"io/ioutil"

	"github.com/adrianriobo/dater/pkg/util/http"
)

func GetFileContent(fileUrl, filePath string) ([]byte, error) {
	if len(fileUrl) > 0 {
		return http.GetFile(fileUrl)
	}
	if len(filePath) == 0 {
		return nil, fmt.Errorf("no source location for file")
	}
	return ioutil.ReadFile(filePath)
}
