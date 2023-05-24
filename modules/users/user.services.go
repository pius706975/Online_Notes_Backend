package users

import (
	"log"
	"os"

	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
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
