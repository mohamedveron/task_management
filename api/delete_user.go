package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id string) {

	res, err := s.svc.DeleteUser(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
