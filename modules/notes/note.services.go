package notes

import (
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
)

type Note_Service struct {
	repo Note_Repo
}

func NewNoteService(repo Note_Repo) *Note_Service {
	return &Note_Service{repo}
}

// ADD NOTE
func (s *Note_Service) AddNewNote(data *models.Note) *libs.Response {
	
	newData, err := s.repo.AddNewNote(data)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}

	return libs.Respond(newData, 200, false)
}