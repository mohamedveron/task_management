package api

import (
	"encoding/json"
	"github.com/mohamedveron/task_management/domains"
	"net/http"
)

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request){

	var newTask NewTask

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := domains.Task{
		ID:           "",
		Name:         newTask.Name,
		Owner:        domains.User{},
		State:        domains.EnumtaskState(newTask.State),
	}

	id, err := s.svc.CreateTask(task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
