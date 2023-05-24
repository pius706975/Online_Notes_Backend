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

// REGISTER
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

// UPDATE
func (r *User_Repo) UpdateUser(data *models.User, ID string) (*models.User, error) {

	result := r.db.Model(data).Where("user_id = ?", ID).Updates(&data).Find(&data).Error
	if result != nil {
		return nil, errors.New("update data failed")
	}

	data.Password = ""
	data.Role = ""

	return data, nil
}

// EMAIL EXISTS
func (r *User_Repo) EmailExists(email string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// USERNAME EXISTS
func (r *User_Repo) UsernameExists(username string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// TOKEN EXISTS
func (r *User_Repo) TokenExists(token string) bool {

	var data models.User

	err := r.db.First(&data, "token_verify = ?", token).Error

	return err == nil
}

//UPDATE TOKEN 
func (r *User_Repo) UpdateToken(ID, token string) error {
	
	var data models.User

	err := r.db.Model(data).Where("user_id = ?", ID).Update("token_verify", token).Error
	if err != nil {
		return errors.New("update token failed")
	}

	return nil
}

// GET USER'S TOKEN
func (r *User_Repo) GetToken(token string) (*models.User, error) {

	var data models.User

	err := r.db.First(&data, "token_verify = ?", token).Error
	if err != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

// GET ALL USERS
func (r *User_Repo) GetAllUsers() (*models.Users, error) {
	
	var data models.Users

	result := r.db.Select("user_id, name, username, email, gender, country, date_of_birth, mobile_number, token_verify, image, created_at, updated_at").Where("role = ?", "user").Order("created_at DESC").Find(&data).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

// GET BY ID
func (r *User_Repo) GetByID(ID string) (*models.User, error) {

	var data models.User

	result := r.db.Select("user_id, name, username, email, gender, country, date_of_birth, mobile_number, token_verify, image, created_at, updated_at").Find(&data, "user_id = ?", ID).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

//GET EMAIL 
func (r *User_Repo) GetEmail(email string) (*models.User, error) {
	
	var data models.User

	result := r.db.First(&data, "email = ?", email).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}
