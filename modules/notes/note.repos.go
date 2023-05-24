package notes

import "gorm.io/gorm"

type Note_Repo struct {
	db *gorm.DB
}

func NewNoteRepo(db *gorm.DB) Note_Repo {
	return Note_Repo{db}
}