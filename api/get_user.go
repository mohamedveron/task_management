package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request, id string) {

	User , err := s.svc.GetUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}