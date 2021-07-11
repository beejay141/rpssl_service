package helpers

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type RestRequestHelper struct {

}

type IRestRequestHelper interface {
	Get(url string, headers map[string]string) (int,[]byte,error)
}


func (service RestRequestHelper) Get(url string, headers map[string]string) (int,[]byte,error) {
	client := resty.New()

	request := client.R()

	if headers != nil {
		request = request.SetHeaders(headers)
	}

	response, err := request.Get(url)
	if response != nil {
		return response.StatusCode(), response.Body(), err
	}

	return 500, nil,errors.New(fmt.Sprintf("Api Service not available: %s",url))
}
