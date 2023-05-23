package users

import (
	"errors"

	"github.com/pius706975/backend/database/models"
	"gorm.io/gorm"
)

type User_Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) User_Repo {
	return User_Repo{db}
}

//REGISTER
func (r *User_Repo) Register(data *models.User) (*models.User, error) {
	
	result := r.db.Create(data).Error
	if result != nil {
		return nil, result
	}

	data.Password = ""
	data.MobileNumber = ""
	data.Role = ""
	data.Image = ""

	return data, nil
}

//EMAIL EXISTS
func (r *User_Repo) EmailExists(email string) (bool, error) {
	
	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

//USERNAME EXISTS
func (r *User_Repo) UsernameExists(username string) (bool, error) {
	
	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

//GET BY ID
func (r *User_Repo) GetByID(ID string) (*models.User, error) {
	
	var data models.User

	result := r.db.Select("user_id, name, username, email, gender, country, date_of_birth, token_verify, is_active, created_at, updated_at").Find(&data, "user_id = ?", ID).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}