package httputils

import "net/http"

// SetupCorsResponse setup response
func SetupCorsResponse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Vary", "Origin")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
}
