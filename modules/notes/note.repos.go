package notes

import (
	"errors"
	"strings"

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

	history := &models.History{
		Note_ID: data.NoteID,
		Status: "Created",
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	err_ := r.db.Create(history).Error
	if err_ != nil {
		return nil, err_
	}

	return data, nil
}

// UPDATE NOTE
func (r *Note_Repo) UpdateNote(data *models.Note, ID string) (*models.Note, error) {
	
	result := r.db.Model(&data).Where("note_id = ?", ID).Updates(&data).Find(&data).Error
	if result != nil {
		return nil, errors.New("update failed")
	}

	return data, nil
}

// DELETE NOTE
func (r *Note_Repo) DeleteNote(ID string) error {
	
	var data models.Note

	result := r.db.Delete(data, "note_id = ?", ID).Error
	if result != nil {
		return result
	}

	return nil
}

// GET BY ID
func (r *Note_Repo) GetByID(ID string) (*models.Note, error) {
	
	var data models.Note

	result := r.db.First(&data, "note_id = ?", ID).Error
	if result != nil {
		return nil, result
	}

	return &data, nil
}

// GET ALL NOTES
func (r *Note_Repo) GetAllNotes() (*models.Notes, error) {
	
	var data models.Notes

	result := r.db.Order("date desc").Find(&data).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	if len(data) <= 0 {
		return nil, errors.New("data is empty")
	}

	return &data, nil
}

// SEARCH NOTE
func (r *Note_Repo) SearchNote(query string) (*models.Notes, error) {
	
	var data models.Notes

	result := r.db.Where("lower(notes.title) ILIKE ?", "%"+strings.ToLower(query)+"%").Find(&data).Error
	if result != nil {
		return nil, result
	}

	if len(data) <= 0 {
		return nil, errors.New("data not found")
	}

	return &data, nil
}