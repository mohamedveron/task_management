package api

import (
	"encoding/json"
	"github.com/mohamedveron/task_management/domains"
	"net/http"
)

func(s *Server) UpdateTask(w http.ResponseWriter, r *http.Request, id string){

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

	res, err := s.svc.UpdateTask(task, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
