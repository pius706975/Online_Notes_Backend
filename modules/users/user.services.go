package users

import (
	"log"
	"os"

	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
	"gorm.io/gorm"
)

type User_Service struct {
	repo User_Repo
}

func NewUserService(repo User_Repo) *User_Service {
	return &User_Service{repo}
}

// REGISTER
func (s *User_Service) Register(data *models.User) *libs.Response {

	emailExists, err := s.repo.EmailExists(data.Email)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}
	if emailExists {
		return libs.Respond("Email already used", 400, true)
	}

	usernameExists, err := s.repo.UsernameExists(data.Username)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}
	if usernameExists {
		return libs.Respond("Username already used", 400, true)
	}

	hashPassword, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}

	data.Password = hashPassword

	tokenVerify, err := libs.CodeCrypt(32)
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}

	data.TokenVerify = tokenVerify

	emailData := libs.EmailData{
		URL:      os.Getenv("BASE_URL") + "/auth/verify_email/" + tokenVerify,
		Username: data.Username,
		Subject:  "Account Verification",
	}

	err = libs.SendEmail(data, &emailData)
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}

	log.Println(err)

	newData, err := s.repo.Register(data)
	if err != nil {
		return libs.Respond(err.Error(), 400, true)
	}

	result, _ := s.repo.GetByID(newData.UserID)

	return libs.Respond(result, 200, false)
}

// UPDATE
func (s *User_Service) UpdateUser(data *models.User, ID string) *libs.Response {

	var user models.User

	err := s.repo.db.Where("user_id = ?", ID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return libs.Respond("Data not found", 404, true)
		} else {
			return libs.Respond(err.Error(), 500, true)
		}
	}

	if data.Password != "" {
		hashPassword, err := libs.HashPassword(data.Password)
		if err != nil {
			return libs.Respond("Password update failed", 400, true)
		}

		data.Password = hashPassword
	}

	emailExists, err := s.repo.EmailExists(data.Email)
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}
	if emailExists {
		return libs.Respond("Email already exists", 400, true)
	}

	usernameExists, err := s.repo.UsernameExists(data.Username)
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}
	if usernameExists {
		return libs.Respond("Username already exists", 400, true)
	}

	if data.Name == "" {
		data.Name = user.Name
	}
	if data.Username == "" {
		data.Username = user.Username
	}
	if data.Email == "" {
		data.Email = user.Email
	}
	if data.Password == "" {
		data.Password = user.Password
	}
	if data.Gender == "" {
		data.Gender = user.Gender
	}
	if data.Country == "" {
		data.Country = user.Country
	}
	if data.DateOfBirth == "" {
		data.DateOfBirth = user.DateOfBirth
	}
	if data.MobileNumber == "" {
		data.MobileNumber = user.MobileNumber
	}
	if data.Role == "" {
		data.Role = user.Role
	}
	if data.Image == "" {
		data.Image = user.Image
	}

	result, err := s.repo.UpdateUser(data, ID)
	if err != nil {
		return libs.Respond(err.Error(), 500, true)
	}

	return libs.Respond(result, 200, false)
}
