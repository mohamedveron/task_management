package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
	"math/rand"
)

func (s *Service) CreateTask(task domains.Task) (string, error){

	// generate new id
	id := randGeneratePassword(4)

	task.ID = id

	s.tasksDB[id] = task

	owner := s.findUserById(task.Owner.ID)

	updatedTask := domains.Task{
		ID:           task.ID,
		Name:         task.Name,
		Owner:       *owner,
		State:        task.State,
	}

	if owner == nil {
		return "not exist", errors.New("User not exist")
	}

	s.tasksDB[id] = updatedTask

	return id, nil
}

func randGeneratePassword(n int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}