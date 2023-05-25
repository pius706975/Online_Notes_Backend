package notes

import (
	"errors"

	"github.com/pius706975/backend/database/models"
	"gorm.io/gorm"
)

type Note_Repo struct {
	db *gorm.DB
}

func NewNoteRepo(db *gorm.DB) Note_Repo {
	return Note_Repo{db}
}

// ADD NOTE
func (r *Note_Repo) AddNewNote(data *models.Note) (*models.Note, error) {
	
	result := r.db.Create(data).Find(&data).Error
	if result != nil {
		return nil, errors.New("create data failed")
	}

	return data, nil
}