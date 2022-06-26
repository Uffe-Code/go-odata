package modelGenerator

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type definition struct {
	name       string
	url        string
	properties []property
}

type property struct {
	name     string
	dataType string
}

type edmxEntitySet struct {
	schema     edmxSchema
	Name       string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}

func (s edmxEntitySet) getEntityType() edmxEntityType {
	entityTypeKey := strings.Replace(s.EntityType, s.schema.Namespace+".", "", 1)
	return s.schema.EntityTypes[entityTypeKey]
}

type edmxProperty struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Nullable string `xml:"Nullable,attr"`
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
	case "Edm.Int64":
		goType = "int64"
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

func (e rawEdmxEntityType) toEdmxEntityType() edmxEntityType {
	entityType := edmxEntityType{
		Name:       e.Name,
		Properties: map[string]edmxProperty{},
	}
	for _, prop := range e.Properties {
		entityType.Properties[prop.Name] = prop
	}
	return entityType
}

type edmxXmlData struct {
	XMLName      xml.Name           `xml:"Edmx"`
	Version      string             `xml:"Version,attr"`
	DataServices []edmxDataServices `xml:"DataServices"`
}

type edmxSchema struct {
	Namespace   string
	EntityTypes map[string]edmxEntityType
	EntitySets  map[string]edmxEntitySet
}

type edmxDataServices struct {
	Schemas []rawEdmxSchema `xml:"Schema"`
}

type rawEdmxSchema struct {
	XMLName     xml.Name            `xml:"Schema"`
	Namespace   string              `xml:"Namespace,attr"`
	EntityTypes []rawEdmxEntityType `xml:"EntityType"`
	Containers  []rawEdmxContainer  `xml:"EntityContainer"`
}

func (s rawEdmxSchema) toSchema() edmxSchema {
	schema := &edmxSchema{
		Namespace:   s.Namespace,
		EntityTypes: map[string]edmxEntityType{},
		EntitySets:  map[string]edmxEntitySet{},
	}
	for _, e := range s.EntityTypes {
		schema.EntityTypes[e.Name] = e.toEdmxEntityType()
	}
	for _, c := range s.Containers {
		for _, s := range c.EntitySets {
			s.schema = *schema
			schema.EntitySets[s.Name] = s
		}
	}
	return *schema
}

type rawEdmxContainer struct {
	EntitySets []edmxEntitySet `xml:"EntitySet"`
}

func parseEdmx(xmlData []byte) (edmxSchema, error) {
	var edmxData edmxXmlData
	err := xml.Unmarshal(xmlData, &edmxData)
	if err != nil {
		return edmxSchema{}, err
	}

	if edmxData.Version != "4.0" {
		return edmxSchema{}, fmt.Errorf("only version 4.0 is supported, got %s", edmxData.Version)
	}

	if len(edmxData.DataServices) != 1 {
		return edmxSchema{}, fmt.Errorf("unexpected amount of <edmx:DataServices> in Edmx source, got %d and expected 1", len(edmxData.DataServices))
	}

	dataServices := edmxData.DataServices[0]
	if len(dataServices.Schemas) != 1 {
		return edmxSchema{}, fmt.Errorf("unexpected amount of Schema in Edmx source, got %d and expected 1", len(dataServices.Schemas))
	}

	return dataServices.Schemas[0].toSchema(), nil
}

func (d *definition) applyEdmxEntitySet(json edmxEntitySet) {
	d.url = json.Name
}

func (d *definition) applyEdmxEntityType(entity edmxEntityType) {
	d.name = entity.Name
	for _, edmxProp := range entity.Properties {
		d.properties = append(d.properties, property{
			name:     edmxProp.Name,
			dataType: edmxProp.goType(),
		})
	}
}
