package modelGenerator

import (
	"io/ioutil"
	"net/http"
)

func fetchEdmx(url string) (edmxDataServices, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return edmxDataServices{}, err
	}
	request.Header.Set("Accept", "application/atom+xml")
	request.Header.Set("DataServiceVersion", "4.0")
	response, err := client.Do(request)
	if err != nil {
		return edmxDataServices{}, err
	}
	defer func() { _ = response.Body.Close() }()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return edmxDataServices{}, err
	}
	return parseEdmx(body)
}
