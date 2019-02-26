package httputils

import (
	"net/http"
)

// WriteJavaScript write js to the response stream
func WriteJavaScript(w http.ResponseWriter, code int, js []byte) error {
	w.Header().Add("Content-Type", "application/x-javascript")
	w.WriteHeader(code)
	w.Write([]byte(js))
	return nil
}
