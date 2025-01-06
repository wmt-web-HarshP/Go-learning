package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody parses the request body into the specified struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
