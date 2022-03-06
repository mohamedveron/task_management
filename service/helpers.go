package service

import "github.com/mohamedveron/task_management/domains"

func (s *Service) findTaskById(id string) *domains.Task{

	if task, ok := s.tasksDB[id]; ok {
		return &task
	}

	return nil
}

func (s *Service) findUserById(id string) *domains.User{

	if User, ok := s.UsersDB[id]; ok {
		return &User
	}

	return nil
}
