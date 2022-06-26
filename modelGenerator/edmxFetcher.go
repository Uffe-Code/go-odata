package modelGenerator

import (
	"io/ioutil"
	"net/http"
)

func fetchEdmx(url string) (edmxSchema, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return edmxSchema{}, err
	}
	response, err := client.Do(request)
	if err != nil {
		return edmxSchema{}, err
	}
	defer func() { _ = response.Body.Close() }()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return edmxSchema{}, err
	}
	return parseEdmx(body)
}
