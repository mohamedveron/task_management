package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) UpdateTask(task domains.Task, id string) (string, error){

	proj := s.findTaskById(id)

	if proj == nil {
		return "not exist", errors.New("task not exist")
	}

	updatedTask := domains.Task{
		ID:           proj.ID,
		Name:         task.Name,
		Owner:        proj.Owner,
		State:        task.State,
	}

    s.tasksDB[id] = updatedTask

	return "updated", nil
}

