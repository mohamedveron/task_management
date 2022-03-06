package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request){

	tasks , err := s.svc.Gettasks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
