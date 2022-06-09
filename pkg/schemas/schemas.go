package schemas

import (
	"fmt"
	"io"
	"os"

	"github.com/a-h/generate"
	"github.com/xuri/xgen"
	"golang.org/x/exp/slices"
)

// //go:generate go run gen-json.go xunit/xunit.json xunit/xunit.go xunit
//go:generate go run gen-xsd.go xunit xunit.xsd xunit xunit

func GenerateFromXSD(inputFolder, xsdFileName, outputFolder, packageName string) error {
	files, err := xgen.GetFileList(inputFolder)
	if err != nil {
		return err
	}
	file := getXSDFile(files, fmt.Sprintf("%s/%s", inputFolder, xsdFileName))
	if file == "" {
		return fmt.Errorf("file %s not found", xsdFileName)
	}

	if err = xgen.NewParser(&xgen.Options{
		FilePath:            file,
		InputDir:            inputFolder,
		OutputDir:           outputFolder,
		Lang:                "Go",
		Package:             packageName,
		IncludeMap:          make(map[string]bool),
		LocalNameNSMap:      make(map[string]string),
		NSSchemaLocationMap: make(map[string]string),
		ParseFileList:       make(map[string]bool),
		ParseFileMap:        make(map[string][]interface{}),
		ProtoTree:           make([]interface{}, 0),
		RemoteSchema:        make(map[string][]byte),
	}).Parse(); err != nil {
		return err
	}
	return nil
}

func GenerateFromJSONSchema(inputFileName, outputFileName, packageName string) error {
	schemas, err := generate.ReadInputFiles(append([]string{}, inputFileName), true)
	if err != nil {
		return err
	}
	g := generate.New(schemas...)
	err = g.CreateTypes()
	if err != nil {
		return err
	}
	var w io.Writer = os.Stdout
	w, err = os.Create(outputFileName)
	if err != nil {
		return err
	}
	generate.Output(w, g, packageName)
	return nil
}

func getXSDFile(files []string, xsdFileName string) string {
	idx := slices.IndexFunc(files,
		func(e string) bool { return e == xsdFileName })
	if idx == -1 {
		return ""
	}
	return files[idx]
}
