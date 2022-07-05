package odataClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type odataDataSet[ModelT any, Def ODataModelDefinition[ModelT]] struct {
	client          *oDataClient
	modelDefinition ODataModelDefinition[ModelT]
}

type ODataDataSet[ModelT any, Def ODataModelDefinition[ModelT]] interface {
	Single(id string) (ModelT, error)
	List(filter ODataFilter) (<-chan ModelT, <-chan error)
	Insert(model ModelT) (ModelT, error)
	Update(id string, model ModelT) (ModelT, error)
	Delete(id string) error

	getCollectionUrl() string
	getSingleUrl(modelId string) string
}

func NewDataSet[ModelT any, Def ODataModelDefinition[ModelT]](client ODataClient, modelDefinition Def) ODataDataSet[ModelT, Def] {
	return odataDataSet[ModelT, Def]{
		client:          client.(*oDataClient),
		modelDefinition: modelDefinition,
	}
}

func (dataSet odataDataSet[ModelT, Def]) getCollectionUrl() string {
	return dataSet.client.baseUrl + dataSet.modelDefinition.Url()
}

func (dataSet odataDataSet[ModelT, Def]) getSingleUrl(modelId string) string {
	return fmt.Sprintf("%s(%s)", dataSet.client.baseUrl+dataSet.modelDefinition.Url(), modelId)
}

type apiSingleResponse[T interface{}] struct {
	Value T `json:"value"`
}

type apiMultiResponse[T interface{}] struct {
	Value    []T    `json:"value"`
	NextLink string `json:"@odata.nextLink"`
}

// ODataFilter represents a OData Filter query
type ODataFilter struct {
	Filter string
}

func (filter ODataFilter) toQueryString() string {
	return fmt.Sprintf("$filter=%s", filter.Filter)
}

// Single model from the API by ID
func (dataSet odataDataSet[ModelT, Def]) Single(id string) (ModelT, error) {
	url := dataSet.getSingleUrl(id)
	request, err := http.NewRequest("GET", url, nil)
	var responseModel ModelT
	if err != nil {
		return responseModel, err
	}
	responseData, err := executeHttpRequest[apiSingleResponse[ModelT]](*dataSet.client, request)
	if err != nil {
		return responseModel, err
	}
	return responseData.Value, nil
}

// List data from the API
func (dataSet odataDataSet[ModelT, Def]) List(filter ODataFilter) (<-chan ModelT, <-chan error) {
	ch := make(chan ModelT)
	errs := make(chan error)

	go func() {
		defer close(ch)
		defer close(errs)

		url := fmt.Sprintf("%s?$top=%d&$skip=0&%s", dataSet.getCollectionUrl(), dataSet.client.defaultPageSize, filter.toQueryString())
		for url != "" {
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				errs <- err
				return
			}
			responseData, err := executeHttpRequest[apiMultiResponse[ModelT]](*dataSet.client, request)
			if err != nil {
				errs <- err
				return
			}

			for _, model := range responseData.Value {
				ch <- model
			}
			if len(responseData.Value) < dataSet.client.defaultPageSize {
				return
			}
			url = responseData.NextLink
		}
	}()

	return ch, errs
}

// Insert a model to the API
func (dataSet odataDataSet[ModelT, Def]) Insert(model ModelT) (ModelT, error) {
	url := dataSet.getCollectionUrl()
	var result ModelT
	jsonData, err := json.Marshal(model)
	if err != nil {
		return result, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return result, err
	}
	request.Header.Set("Content-Type", "application/json;odata.metadata=minimal")
	request.Header.Set("Prefer", "return=representation")
	return executeHttpRequest[ModelT](*dataSet.client, request)
}

// Update a model in the API
func (dataSet odataDataSet[ModelT, Def]) Update(id string, model ModelT) (ModelT, error) {
	url := dataSet.getSingleUrl(id)
	var result ModelT
	jsonData, err := json.Marshal(model)
	if err != nil {
		return result, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return result, err
	}
	request.Header.Set("Content-Type", "application/json;odata.metadata=minimal")
	request.Header.Set("Prefer", "return=representation")
	return executeHttpRequest[ModelT](*dataSet.client, request)
}

// Delete a model from the API
func (dataSet odataDataSet[ModelT, Def]) Delete(id string) error {
	url := dataSet.getSingleUrl(id)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	dataSet.client.mapHeadersToRequest(request)
	response, err := dataSet.client.httpClient.Do(request)
	if err != nil {
		return err
	}
	_ = response.Body.Close()
	return nil
}
