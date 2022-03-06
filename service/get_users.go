package service

import "github.com/mohamedveron/task_management/domains"

func (s *Service) GetUsers() ([]domains.User, error){

	UsersList := []domains.User{}

	for _, Users := range s.UsersDB {

		UsersList = append(UsersList, Users)
	}

	return UsersList, nil
}
