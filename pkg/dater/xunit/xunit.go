package xunit

import (
	"encoding/xml"

	"github.com/adrianriobo/dater/pkg/util/file"
	"github.com/adrianriobo/dater/pkg/util/logging"

	xunitSchema "github.com/adrianriobo/dater/pkg/schemas/xunit"
)

const (
	statusSuccess string = "SUCCESS"
	statusFail    string = "FAIL"
)

func GlobalStatusRemote(xunitUrl, xunitPath string) (string, error) {
	xunitContent, err := file.GetFileContent(xunitUrl, xunitPath)
	if err != nil {
		return "", err
	}
	return GlobalStatus(xunitContent)
}

func GlobalStatus(xunitContent []byte) (string, error) {
	xunitContentAsStruct := xunitSchema.Testsuites{}
	err := xml.Unmarshal(xunitContent, &xunitContentAsStruct)
	if err != nil {
		return "", err
	}
	logging.Debugf("Analyzed xunit file got %s failures", xunitContentAsStruct.FailuresAttr)
	if xunitContentAsStruct.FailuresAttr == "0" {
		return statusSuccess, nil
	}
	return statusFail, nil
}
