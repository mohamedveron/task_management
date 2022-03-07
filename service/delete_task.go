package service

func (s *Service) DeleteTask(id string) (string, error){

	if _, ok := s.tasksDB[id]; ok {
		return "not exist", nil
	}

	delete(s.tasksDB, id)

	return "deleted", nil
}

