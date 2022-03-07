package api

import (
	"encoding/json"
	"github.com/mohamedveron/task_management/domains"
	"net/http"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := domains.User{
		ID:        "",
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	id, err := s.svc.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
