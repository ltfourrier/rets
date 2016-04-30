package rets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// UnmarshalHTTPBody takes a http.Request and returns the unmarshalled
// representation of its JSON body.
func UnmarshalHTTPBody(request *http.Request, v interface{}) error {
	payload, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(payload, v)
	return err
}
