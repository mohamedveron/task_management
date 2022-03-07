package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) CreateTask(task domains.Task) (string, error) {

	// generate new id
	id := RandGeneratePassword(16)

	task.ID = id

	s.tasksDB[id] = task

	owner, err := s.GetUserById(task.Owner.ID)

	if owner == nil || err != nil {
		return "not exist", errors.New("User not exist")
	}

	newTask := domains.Task{
		ID:             task.ID,
		Name:           task.Name,
		Owner:          *owner,
		Estimation:     task.Estimation,
		ReminderPeriod: task.ReminderPeriod,
		State:          task.State,
	}

	s.tasksDB[id] = newTask

	return id, nil
}


