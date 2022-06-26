package modelGenerator

import (
	"fmt"
)

type generator struct {
	apiUrl string
}

func (g generator) metadataUrl() string {
	return g.apiUrl + "$metadata"
}

func generateModelStruct(set edmxEntitySet) string {
	entityType := set.getEntityType()

	structString := fmt.Sprintf("type %s struct {", entityType.Name)

	propertyKeys := sortedKeys(entityType.Properties)

	for _, propertyKey := range propertyKeys {
		prop := entityType.Properties[propertyKey]
		structString += fmt.Sprintf("\n\t%s %s", prop.Name, prop.goType())
	}

	return structString + "\n}"
}

func generateModelDefinition(set edmxEntitySet) string {
	entityType := set.getEntityType()

	// type modelDefinition struct { name string, url string }

	return fmt.Sprintf(`func %sDefinition() odataClient.ODataModelCollection[%s] {
	return modelDefinition{name: "%s", url: "%s"}
}`, entityType.Name, entityType.Name, entityType.Name, set.Name)
}

func generateCodeFromSchema(packageName string, schema edmxSchema) string {
	goCode := fmt.Sprintf(`package %s

import "github.com/Uffe-Code/go-nullable"

`, packageName)

	for _, set := range schema.EntitySets {
		goCode += "\n" + generateModelStruct(set) + "\n"
		goCode += "\n" + generateModelDefinition(set) + "\n"
	}

	return goCode
}
