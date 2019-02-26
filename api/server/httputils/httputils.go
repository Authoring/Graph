package httputils

import (
	"context"
	"net/http"
)

// APIFunc is a type for the handler
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
