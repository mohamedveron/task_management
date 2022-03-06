package service

import (
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) Gettasks() ([]domains.Task, error){

	tasksList := []domains.Task{}

	for _, task := range s.tasksDB {

		tasksList = append(tasksList, task)
	}

	return tasksList, nil
}
