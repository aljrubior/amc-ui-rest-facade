package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	errors2 "github.com/aljrubior/amc-ui-rest-facade/errors"
	"io/ioutil"
	"net/http"
)

type BaseHttpClient struct {
}

func (_ *BaseHttpClient) ThrowError(response *http.Response) error {

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errors.New(response.Status)
	}

	var body map[string]string

	if err := json.Unmarshal(data, &body); err != nil {
		errors.New(response.Status)
	}

	switch response.StatusCode {
	case 401:
		return errors2.NewUnauthorizedError(response)
	default:
		return errors.New(fmt.Sprintf("%s. Reason: %s", response.Status, string(data)))
	}
}
