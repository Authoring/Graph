package httputils

import (
	"io/ioutil"
	"net/http"
)

// ReadBody read body
func ReadBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}
