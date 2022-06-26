package modelGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	ApiUrl        string
	DirectoryPath string
}

func (g Generator) metadataUrl() string {
	return strings.TrimRight(g.ApiUrl, "/") + "/$metadata"
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

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(code)
	return err
}
