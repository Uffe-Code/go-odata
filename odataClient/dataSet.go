package odataClient

import (
	"fmt"
	"net/http"
)

type odataDataSet[ModelT any, Def ODataModelDefinition[ModelT]] struct {
	client          *oDataClient
	modelDefinition ODataModelDefinition[ModelT]
}

type ODataDataSet[ModelT any, Def ODataModelDefinition[ModelT]] interface {
	Single(id int) (ModelT, error)
	List(filter ODataFilter) (<-chan ModelT, <-chan error)

	getCollectionUrl() string
	getSingleUrl(modelId int) string
}

func newDataSet[ModelT any, Def ODataModelDefinition[ModelT]](client ODataClient, modelDefinition Def) ODataDataSet[ModelT, Def] {
	return odataDataSet[ModelT, Def]{
		client:          client.(*oDataClient),
		modelDefinition: modelDefinition,
	}
}

func (dataSet odataDataSet[ModelT, Def]) getCollectionUrl() string {
	return dataSet.client.baseUrl + dataSet.modelDefinition.Url()
}

func (dataSet odataDataSet[ModelT, Def]) getSingleUrl(modelId int) string {
	return fmt.Sprintf("%s(%d)", dataSet.client.baseUrl+dataSet.modelDefinition.Url(), modelId)
}

type apiSingleResponse[T interface{}] struct {
	Value T `json:"value"`
}

type apiMultiResponse[T interface{}] struct {
	Value    []T    `json:"value"`
	NextLink string `json:"@odata.nextLink"`
}

type ODataFilter struct {
	filter string
}

func (filter ODataFilter) toQueryString() string {
	return fmt.Sprintf("$filter=%s", filter.filter)
}

func (dataSet odataDataSet[ModelT, Def]) Single(id int) (ModelT, error) {
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
