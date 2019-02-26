package status

import (
	"context"
	"net/http"

	"github.com/Authoring/Graph/api/server/httputils"
)

type statusResponse struct {
	Status bool `json:"status"`
}

func (s *statusRouter) getStatus(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	status, err := s.backend.Status()

	if err != nil {
		return err
	}

	var code int
	if status {
		code = http.StatusOK
	} else {
		code = http.StatusBadRequest
	}

	res := statusResponse{
		Status: status,
	}

	return httputils.WriteJSON(w, code, res)
}
