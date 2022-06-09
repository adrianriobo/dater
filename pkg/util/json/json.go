package json

import (
	"bytes"

	xj "github.com/basgys/goxml2json"
)

func ConvertToJSON(xunitFileContent []byte) (json string, err error) {
	xunitContent := bytes.NewReader(xunitFileContent)
	converted, err := xj.Convert(xunitContent)
	if err != nil {
		return "", err
	}
	return converted.String(), nil
}
