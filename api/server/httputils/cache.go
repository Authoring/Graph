package httputils

import "net/http"

// DisableCache set cache response headers
func DisableCache(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Max-Age", "1")
	(*w).Header().Set("Cache-Control", "public, max-age=1")
}
