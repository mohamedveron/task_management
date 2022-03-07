package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetTask(w http.ResponseWriter, r *http.Request, id string) {
	task , err := s.svc.GetTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
