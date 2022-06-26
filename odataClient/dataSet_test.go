package odataClient

import (
	"encoding/json"
	"github.com/Uffe-Code/go-nullable/nullable"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testModel struct {
	Id          int
	Number      string
	Name        string
	ParentId    nullable.Nullable[int]
	Description nullable.Nullable[string]
}

type testModelDefinition[T any] struct {
}

func (t testModelDefinition[T]) DataSet(client ODataClient) ODataDataSet[T, ODataModelDefinition[T]] {
	return newDataSet[T](client, t)
}

func (t testModelDefinition[T]) Name() string {
	return "Person"
}

func (t testModelDefinition[T]) Url() string {
	return "People"
}

func newTestModelDefinition() ODataModelCollection[testModel] {
	return testModelDefinition[testModel]{}
}

func TestNewDataSet(t *testing.T) {
	client := New("http://test.api/")
	dataSet := newTestModelDefinition().DataSet(client)
	assert.Equal(t, "http://test.api/People", dataSet.getCollectionUrl())
	assert.Equal(t, "http://test.api/People(5)", dataSet.getSingleUrl(5))
}

func TestNewDataSet_WithoutSlash(t *testing.T) {
	client := New("http://test.api")
	dataSet := newTestModelDefinition().DataSet(client)
	assert.Equal(t, "http://test.api/People", dataSet.getCollectionUrl())
	assert.Equal(t, "http://test.api/People(5)", dataSet.getSingleUrl(5))
}

func TestOdataDataSet_Single(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/People(5)" {
			writer.WriteHeader(404)
			return
		}
		writer.WriteHeader(200)
		data, _ := json.Marshal(struct {
			Value testModel `json:"value"`
		}{
			Value: testModel{
				Id:          5,
				Number:      "002",
				Name:        "Donald Duck",
				ParentId:    nullable.Null[int](),
				Description: nullable.Value("Test description"),
			},
		})
		_, _ = writer.Write(data)
	}))
	defer testServer.Close()

	client := New(testServer.URL)
	def := newTestModelDefinition()
	dataSet := def.DataSet(client)
	model, err := dataSet.Single(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, model.Id)
	assert.Equal(t, "002", model.Number)
	assert.Equal(t, "Donald Duck", model.Name)
	assert.False(t, model.ParentId.IsValid)
	assert.True(t, model.Description.IsValid)
	assert.Equal(t, "Test description", model.Description.Data)
}

func TestOdataDataSet_List(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/People" {
			writer.WriteHeader(404)
			return
		}
		writer.WriteHeader(200)
		data, _ := json.Marshal(struct {
			Value []testModel `json:"value"`
		}{
			Value: []testModel{
				{
					Id:          3,
					Number:      "001",
					Name:        "Uffe Code",
					ParentId:    nullable.Null[int](),
					Description: nullable.Value("Test description"),
				},
				{
					Id:          5,
					Number:      "002",
					Name:        "Donald Duck",
					ParentId:    nullable.Null[int](),
					Description: nullable.Value("Test description"),
				},
			},
		})
		_, _ = writer.Write(data)
	}))
	defer testServer.Close()

	client := New(testServer.URL)
	def := newTestModelDefinition()
	dataSet := def.DataSet(client)
	models, _ := dataSet.List(ODataFilter{})

	i := 0
	for model := range models {
		if i == 0 {
			assert.Equal(t, 3, model.Id)
			assert.Equal(t, "001", model.Number)
		} else {
			assert.Equal(t, 5, model.Id)
			assert.Equal(t, "002", model.Number)
		}
		assert.False(t, model.ParentId.IsValid)
		assert.True(t, model.Description.IsValid)
		assert.Equal(t, "Test description", model.Description.Data)
		i++
	}
}
