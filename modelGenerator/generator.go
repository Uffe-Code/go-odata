package modelGenerator

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Generator struct {
	ApiUrl        string
	DirectoryPath string
}

func (g Generator) metadataUrl() string {
	return g.ApiUrl + "$metadata"
}

func (g Generator) GenerateCode() error {
	dirPath, err := filepath.Abs(g.DirectoryPath)
	if err != nil {
		return err
	}

	edmx, err := fetchEdmx(g.metadataUrl())
	if err != nil {
		return err
	}

	packageName := filepath.Base(dirPath)
	code := generateCodeFromSchema(packageName, edmx)
	filePath := fmt.Sprintf("%s%s%s", dirPath, string(filepath.Separator), "modelDefinitions.go")

	return ioutil.WriteFile(filePath, []byte(code), 0777)
}
