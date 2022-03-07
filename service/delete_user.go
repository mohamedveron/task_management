package service

func (s *Service) DeleteUser(id string) (string, error){

	if _, ok := s.UsersDB[id]; ok {
		return "not exist", nil
	}

	delete(s.UsersDB, id)

	return "deleted", nil
}
