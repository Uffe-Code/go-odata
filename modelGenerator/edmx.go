package modelGenerator

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type rawEdmxEntitySet struct {
	Name       string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}

func (es rawEdmxEntitySet) toEntitySet(schema edmxSchema) edmxEntitySet {
	return edmxEntitySet{
		schema:     schema,
		Name:       es.Name,
		EntityType: es.EntityType,
	}
}

type edmxEntitySet struct {
	schema     edmxSchema
	Name       string
	EntityType string
}

func (s edmxEntitySet) getEntityType() edmxEntityType {
	namespace := s.EntityType[0:strings.LastIndex(s.EntityType, ".")]
	entityTypeKey := s.EntityType[len(namespace)+1 : len(s.EntityType)]
	return s.schema.dataService.Schemas[namespace].EntityTypes[entityTypeKey]
}

type edmxProperty struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Nullable string `xml:"Nullable,attr"`
	schema   edmxSchema
}

func (p edmxProperty) goType() string {
	propertyType := p.Type
	isCollection := false
	if strings.HasPrefix(p.Type, "Collection(") {
		isCollection = true
		propertyType = p.Type[11 : len(p.Type)-1]
	}
	goType := "interface{}"
	switch propertyType {
	case "Edm.String":
		goType = "string"
	case "Edm.Int32":
		goType = "int32"
	case "Edm.Int64":
		goType = "int64"
	case "Edm.Double":
		goType = "float64"
	case "Edm.Boolean":
		goType = "bool"
	case "Edm.DateTime", "Edm.DateTimeOffset", "Edm.Date": // todo: consider using a separate data type for dates without time
		goType = "time.Time"
	default:
		if strings.HasPrefix(propertyType, p.schema.Namespace) {
			entityTypeKey := strings.Replace(propertyType, p.schema.Namespace+".", "", 1)
			if enumType, ok := p.schema.EnumTypes[entityTypeKey]; ok {
				goType = enumType.Name
			}
			if complexType, ok := p.schema.ComplexTypes[entityTypeKey]; ok {
				goType = complexType.Name
			}
		}
	}

	if !isCollection && (p.Nullable == "" || strings.ToLower(p.Nullable) == "true") {
		goType = "nullable.Nullable[" + goType + "]"
	}

	if isCollection {
		goType = "[]" + goType
	}
	return goType
}

type edmxEntityType struct {
	Name       string
	Properties map[string]edmxProperty
}

type rawEdmxEntityType struct {
	Name       string         `xml:"Name,attr"`
	Properties []edmxProperty `xml:"Property"`
}

func (e rawEdmxEntityType) toEdmxEntityType(schema edmxSchema) edmxEntityType {
	entityType := edmxEntityType{
		Name:       e.Name,
		Properties: map[string]edmxProperty{},
	}
	for _, prop := range e.Properties {
		prop.schema = schema
		entityType.Properties[prop.Name] = prop
	}
	return entityType
}

type edmxXmlData struct {
	XMLName      xml.Name              `xml:"Edmx"`
	Version      string                `xml:"Version,attr"`
	DataServices []rawEdmxDataServices `xml:"DataServices"`
}

type edmxSchema struct {
	dataService  edmxDataServices
	Namespace    string
	EntityTypes  map[string]edmxEntityType
	EntitySets   map[string]edmxEntitySet
	EnumTypes    map[string]edmxEnumType
	ComplexTypes map[string]edmxEntityType
}

type rawEdmxDataServices struct {
	Schemas []rawEdmxSchema `xml:"Schema"`
}

func (ds *rawEdmxDataServices) toDataService() edmxDataServices {
	dataService := &edmxDataServices{Schemas: map[string]edmxSchema{}}
	for _, s := range ds.Schemas {
		sc := s.toSchema(*dataService)
		dataService.Schemas[sc.Namespace] = sc
	}
	return *dataService
}

type edmxDataServices struct {
	Schemas map[string]edmxSchema
}

type rawEdmxSchema struct {
	XMLName      xml.Name            `xml:"Schema"`
	Namespace    string              `xml:"Namespace,attr"`
	EntityTypes  []rawEdmxEntityType `xml:"EntityType"`
	Containers   []rawEdmxContainer  `xml:"EntityContainer"`
	EnumTypes    []edmxEnumType      `xml:"EnumType"`
	ComplexTypes []rawEdmxEntityType `xml:"ComplexType"`
}

func (s rawEdmxSchema) toSchema(services edmxDataServices) edmxSchema {
	schema := &edmxSchema{
		dataService:  services,
		Namespace:    s.Namespace,
		EntityTypes:  map[string]edmxEntityType{},
		EntitySets:   map[string]edmxEntitySet{},
		EnumTypes:    map[string]edmxEnumType{},
		ComplexTypes: map[string]edmxEntityType{},
	}
	for _, e := range s.EntityTypes {
		schema.EntityTypes[e.Name] = e.toEdmxEntityType(*schema)
	}
	for _, c := range s.Containers {
		for _, es := range c.EntitySets {
			entitySet := es.toEntitySet(*schema)
			schema.EntitySets[entitySet.Name] = entitySet
		}
	}
	for _, enum := range s.EnumTypes {
		schema.EnumTypes[enum.Name] = enum
	}
	for _, complexType := range s.ComplexTypes {
		schema.ComplexTypes[complexType.Name] = complexType.toEdmxEntityType(*schema)
	}
	return *schema
}

type rawEdmxContainer struct {
	EntitySets []rawEdmxEntitySet `xml:"EntitySet"`
}

type edmxEnumType struct {
	XMLName xml.Name         `xml:"EnumType"`
	Name    string           `xml:"Name,attr"`
	Members []edmxEnumMember `xml:"Member"`
}

type edmxEnumMember struct {
	XMLName xml.Name `xml:"Member"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:"Value,attr"`
}

type apiErrorMessage struct {
	Message string `xml:"message"`
}

func parseEdmx(xmlData []byte) (edmxDataServices, error) {
	var edmxData edmxXmlData
	err := xml.Unmarshal(xmlData, &edmxData)
	if err != nil {
		var apiErr apiErrorMessage
		err2 := xml.Unmarshal(xmlData, &apiErr)
		if err2 == nil {
			return edmxDataServices{}, fmt.Errorf("error from API: %s", apiErr.Message)
		}
		return edmxDataServices{}, err
	}

	if edmxData.Version != "4.0" {
		return edmxDataServices{}, fmt.Errorf("only version 4.0 is supported, got %s", edmxData.Version)
	}

	if len(edmxData.DataServices) != 1 {
		return edmxDataServices{}, fmt.Errorf("unexpected amount of <edmx:DataServices> in Edmx source, got %d and expected 1", len(edmxData.DataServices))
	}

	dataServices := edmxData.DataServices[0]
	return dataServices.toDataService(), nil
}
