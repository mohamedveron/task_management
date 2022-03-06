package service

import "github.com/mohamedveron/task_management/domains"

type Service struct {
	tasksDB	map[string]domains.Task
	UsersDB map[string]domains.User
}

func NewService(tasksDB	map[string]domains.Task, UsersDB map[string]domains.User) *Service {

	return &Service{
		tasksDB: tasksDB,
		UsersDB: UsersDB,
	}
}
