package history

import "github.com/pius706975/backend/libs"

type History_Service struct {
	repo History_Repo
}

func NewHistoryService(repo History_Repo) History_Service {
	return History_Service{repo}
}

func (s *History_Service) GetAllHistories() *libs.Response {
	
	data, err := s.repo.GetAllHistories()
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}

	return libs.Respond(data, 200, false)
}