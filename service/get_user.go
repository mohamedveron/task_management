package service

import (
	"errors"
	"github.com/mohamedveron/task_management/domains"
)

func (s *Service) GetUserById(id string) (*domains.User, error){

	if User, ok := s.UsersDB[id]; ok {
		return &User, nil
	}

	return nil, errors.New("user not exist")
}
