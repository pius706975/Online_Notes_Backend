package history

import (
	"errors"

	"github.com/pius706975/backend/database/models"
	"gorm.io/gorm"
)

type History_Repo struct {
	db *gorm.DB
}

func NewHistoryRepo(db *gorm.DB) History_Repo {
	return History_Repo{db}
}

func (r *History_Repo) GetAllHistories() (*models.Histories, error) {
	
	var data models.Histories

	result := r.db.Find(&data).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	if len(data) <= 0 {
		return nil, errors.New("data is empty")
	}

	return &data, nil
}