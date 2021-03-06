package odataClient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type oDataClient struct {
	baseUrl         string
	headers         map[string]string
	httpClient      *http.Client
	defaultPageSize int
}

// ODataClient represents a connection to the OData REST API
type ODataClient interface {
	Wrapper
	AddHeader(key string, value string)
}

// Wrapper represents a wrapper around the OData client if you have build own code around the OData itself, for authentication etc
type Wrapper interface {
	ODataClient() ODataClient
}

func New(baseUrl string) ODataClient {
	client := &oDataClient{
		baseUrl: strings.TrimRight(baseUrl, "/") + "/",
		headers: map[string]string{
			"DataServiceVersion": "4.0",
			"OData-Version":      "4.0",
			"Accept":             "application/json",
		},
		defaultPageSize: 1000,
	}

	httpTransport := &http.Transport{}
	client.httpClient = &http.Client{
		Transport: httpTransport,
	}

	return client
}

// AddHeader will add a custom HTTP Header to the API requests
func (client *oDataClient) AddHeader(key string, value string) {
	client.headers[strings.ToLower(key)] = value
}

// ODataClient will return self, so it also works as a wrapper in case we don't have a wrapper
func (client *oDataClient) ODataClient() ODataClient {
	return client
}

func (client oDataClient) mapHeadersToRequest(req *http.Request) {
	for key, value := range client.headers {
		req.Header.Set(key, value)
	}
}

func executeHttpRequest[T interface{}](client oDataClient, req *http.Request) (T, error) {
	client.mapHeadersToRequest(req)
	response, err := client.httpClient.Do(req)
	var responseData T
	if err != nil {
		return responseData, err
	}
	defer func() { _ = response.Body.Close() }()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseData, err
	}
	err = json.Unmarshal(body, &responseData)
	return responseData, err
}
