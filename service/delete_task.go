package service

func (s *Service) DeleteTask(id string) (string, error){

	delete(s.tasksDB, id)

	return "deleted", nil
}

