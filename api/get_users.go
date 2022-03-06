package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request){
	Users , err := s.svc.GetUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Users)
}
