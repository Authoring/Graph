package httputils

import (
	"fmt"
	"net/http"
)

// WriteError write an error out to the client
func WriteError(w http.ResponseWriter, code int, err error) {
	http.Error(w, fmt.Sprintf("%v", err), code)
}
