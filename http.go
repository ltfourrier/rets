package rets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalHTTPBody(request *http.Request, v interface{}) error {
	payload, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(payload, v)
	return err
}
