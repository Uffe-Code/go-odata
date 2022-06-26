package modelGenerator

import (
	"fmt"
	"strconv"
)

func generateModelStruct(entityType edmxEntityType) string {
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

	return fmt.Sprintf(`func %sDefinition() odataClient.ODataModelCollection[%s] {
	return modelDefinition[%s]{name: "%s", url: "%s"}
}`, entityType.Name, entityType.Name, entityType.Name, entityType.Name, set.Name)
}

func generateEnumStruct(enum edmxEnumType) string {
	stringValues := map[string]string{}
	intValues := map[int64]string{}
	isIntValues := true

	for _, member := range enum.Members {
		stringValues[member.Name] = member.Value
		i, err := strconv.ParseInt(member.Value, 10, 64)
		if err != nil {
			isIntValues = false
		} else {
			intValues[i] = member.Name
		}
	}

	goType := "string"
	if isIntValues {
		goType = "int64"
	}
	goString := fmt.Sprintf(`type %s %s

const (`, enum.Name, goType)

	if isIntValues {
		intKeys := sortedKeys(intValues)
		for _, i := range intKeys {
			key := intValues[i]
			goString += fmt.Sprintf("\n\t%s %s = %d", key, enum.Name, i)
		}
	} else {
		stringKeys := sortedKeys(stringValues)
		for _, key := range stringKeys {
			str := stringValues[key]
			goString += fmt.Sprintf("\n\t%s %s = \"%s\"", key, enum.Name, str)
		}
	}

	return goString + "\n)"
}

func generateCodeFromSchema(packageName string, dataService edmxDataServices) string {
	goCode := fmt.Sprintf(`package %s

import "github.com/Uffe-Code/go-nullable/nullable"

type modelDefinition[T any] struct { name string; url string }

func (md modelDefinition[T]) Name() string {
	return md.name
}

func (md modelDefinition[T]) Url() string {
	return md.url
}

func (md modelDefinition[T]) DataSet(client odataClient.ODataClient) odataClient.ODataDataSet[T, odataClient.ODataModelDefinition[T]] {
	return odataClient.NewDataSet[T](client, md)
}

`, packageName)

	for _, schema := range dataService.Schemas {
		for _, enum := range schema.EnumTypes {
			goCode += "\n" + generateEnumStruct(enum) + "\n"
		}

		for _, complexType := range schema.ComplexTypes {
			goCode += "\n" + generateModelStruct(complexType) + "\n"
		}

		for _, set := range schema.EntitySets {
			goCode += "\n" + generateModelStruct(set.getEntityType()) + "\n"
			goCode += "\n" + generateModelDefinition(set) + "\n"
		}
	}

	return goCode
}
