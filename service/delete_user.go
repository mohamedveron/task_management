package service

func (s *Service) DeleteUser(id string) (string, error){

	delete(s.UsersDB, id)

	return "deleted", nil
}
