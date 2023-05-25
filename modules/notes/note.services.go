package notes

import (
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
	"gorm.io/gorm"
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

// DELETE NOTE
func (s *Note_Service) DeleteNote(ID string) *libs.Response {
	
	_, err := s.repo.GetByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return libs.Respond("Data not found", 404, true)
		} else {
			return libs.Respond(err.Error(), 500, true)
		}
	}

	err = s.repo.DeleteNote(ID)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}

	result := map[string]string{"message": "Note has been deleted"}

	return libs.Respond(result, 200, false)
}

// GET ALL NOTES
func (s *Note_Service) GetAllNotes() *libs.Response {
	
	data, err := s.repo.GetAllNotes()
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}

	return libs.Respond(data, 200, false)
}