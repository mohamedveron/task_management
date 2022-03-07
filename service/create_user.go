package service

import "github.com/mohamedveron/task_management/domains"

func (s *Service) CreateUser(user domains.User) (string, error) {

	// generate new id
	id := RandGeneratePassword(4)

	user = domains.User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	s.UsersDB[id] = user

	return id, nil
}
