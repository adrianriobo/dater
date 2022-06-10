package schemas

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/generate"
	"github.com/adrianriobo/dater/pkg/util/logging"
	"github.com/xuri/xgen"
)

// //go:generate go run gen-json.go xunit/xunit.json xunit/xunit.go xunit
//go:generate go run gen-xsd.go xunit xunit xunit

func StructFromXSD(sourceFolder, outputFolder, packageName string) error {
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			if err := StructFromXSD(
				fmt.Sprintf("%s/%s", sourceFolder, fileName),
				outputFolder,
				packageName); err != nil {
				return err
			}
		} else {
			if filepath.Ext(fileName) != ".xsd" {
				logging.Errorf("%s extension is not supported", filepath.Ext(fileName))
				break
			}
			if err = xgen.NewParser(&xgen.Options{
				FilePath:            fmt.Sprintf("%s/%s", sourceFolder, fileName),
				InputDir:            sourceFolder,
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
		}
	}
	return nil
}

func StructFromJSONSchema(sourceFolder, outputFolder, packageName string) error {
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			if err := StructFromJSONSchema(
				fmt.Sprintf("%s/%s", sourceFolder, fileName),
				outputFolder,
				packageName); err != nil {
				return err
			}
		} else {
			switch ext := filepath.Ext(fileName); ext {
			case ".yaml":
				// https://github.com/kubernetes-sigs/yaml
				fmt.Println("OS X.")
			case ".json":
				targetFileName := fmt.Sprintf("%s.go", strings.TrimSuffix(fileName, ext))
				if err := generateFromJSONSchema(
					file.Name(),
					targetFileName,
					packageName); err != nil {
					return nil
				}
			default:
				logging.Errorf("%s extension is not supported", ext)
			}
		}
	}
	return nil
}

func generateFromJSONSchema(inputFileName, outputFileName, packageName string) error {
	schemas, err := generate.ReadInputFiles(append([]string{}, inputFileName), true)
	if err != nil {
		return err
	}
	g := generate.New(schemas...)
	err = g.CreateTypes()
	if err != nil {
		return err
	}
	writer, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	generate.Output(writer, g, packageName)
	return nil
}
