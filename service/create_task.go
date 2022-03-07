package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) CreateTask(task domains.Task) (string, error) {

	// generate new id
	id := RandGeneratePassword(16)

	owner, err := s.GetUserById(task.Owner.ID)

	// check overlap
	for _, v := range s.tasksDB {

		if v.Owner.ID == task.Owner.ID && v.State == domains.EnumtaskStateInProgress{
			return "overlap task detected", errors.New("overlap task")
		}
	}

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


