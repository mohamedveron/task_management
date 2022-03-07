package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) GetTask(taskId string) (*domains.Task, error){

	if val, ok := s.tasksDB[taskId]; ok {
		return &val, nil
	}

	return nil, errors.New("task not exist")
}
