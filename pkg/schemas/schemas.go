package schemas

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/generate"
	"github.com/adrianriobo/dater/pkg/util/logging"
	"github.com/xuri/xgen"
	"sigs.k8s.io/yaml"
)

// //go:generate go run gen-json.go fedora-ci fedora-ci fedoraci
//go:generate go run gen-xsd.go xunit xunit xunit

func StructFromXSD(sourceFolder, outputFolder, packageName string) error {
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}
	fileNames, dirNames := adaptFiles(files)
	for _, dirName := range dirNames {
		if err := StructFromXSD(
			fmt.Sprintf("%s/%s", sourceFolder, dirName),
			outputFolder,
			packageName); err != nil {
			return err
		}
	}
	for _, fileName := range fileNames {
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
	return nil
}

func StructFromJSONSchema(sourceFolder, outputFolder, packageName string) error {
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}
	fileNames, dirNames := adaptFiles(files)
	for _, dirName := range dirNames {
		if err := StructFromJSONSchema(
			fmt.Sprintf("%s/%s", sourceFolder, dirName),
			outputFolder,
			packageName); err != nil {
			return err
		}
	}
	var yamlFiles []string
	var jsonFiles []string
	for _, fileName := range fileNames {
		switch ext := filepath.Ext(fileName); ext {
		case ".yaml":
			yamlFiles = append(yamlFiles, fileName)
		case ".json":
			jsonFiles = append(jsonFiles, fileName)
		default:
			logging.Errorf("%s extension is not supported", ext)
		}
	}
	if len(yamlFiles) > 0 {
		yamlToJsonFiles, err := jsonSchemaFromYAML(
			yamlFiles,
			sourceFolder,
			fmt.Sprintf("%s/%s", sourceFolder, "generated"))
		if err != nil {
			return err
		}
		jsonFiles = append(jsonFiles, yamlToJsonFiles...)
	}
	if len(jsonFiles) > 0 {
		return generateStructFromJSONSchema(jsonFiles, outputFolder, packageName)
	}
	return nil
}

func generateStructFromJSONSchema(inputFilesNames []string, outputFolder, packageName string) error {
	outputFileName := fmt.Sprintf("%s/%s.go", outputFolder, packageName)
	schemas, err := generate.ReadInputFiles(inputFilesNames, false) // append([]string{}, inputFileName)
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

func jsonSchemaFromYAML(yamlFileNames []string, sourceFolder, targetFolder string) ([]string, error) {
	var jsonFiles []string
	// Check if targetFolder exists
	err := os.MkdirAll(targetFolder, os.ModePerm)
	if err != nil {
		return nil, err
	}
	for _, yamlFileName := range yamlFileNames {
		yamlFileContent, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", sourceFolder, yamlFileName))
		if err != nil {
			return nil, err
		}
		jsonSchemaContent, err := yaml.YAMLToJSON(yamlFileContent)
		if err != nil {
			return nil, err
		}
		// Need to edit $id as asks for uri need to check if random is fine
		// Move initial id ...to title
		// Need to add title to be used as struct type name
		ext := filepath.Ext(yamlFileName)
		targetFile := fmt.Sprintf("%s/%s.json",
			targetFolder,
			strings.TrimSuffix(yamlFileName, ext))
		jsonSchemaTempFile, err := os.Create(targetFile)
		if err != nil {
			return nil, err
		}
		defer jsonSchemaTempFile.Close()
		_, err = jsonSchemaTempFile.Write(jsonSchemaContent)
		if err != nil {
			return nil, err
		}
		jsonFiles = append(jsonFiles, targetFile)
	}
	return jsonFiles, nil
}

func adaptFiles(source []fs.FileInfo) (fileNames, dirNames []string) {
	for _, file := range source {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return
}
