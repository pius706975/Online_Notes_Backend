package auth

import (

	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
	"github.com/pius706975/backend/modules/users"
)

type Auth_Service struct {
	repo users.User_Repo
}

type tokenRes struct {
	Token string `json:"token"`
}

func NewAuthService(repo users.User_Repo) *Auth_Service {
	return &Auth_Service{repo}
}

//LOGIN
func (s *Auth_Service) Login(body *models.User) *libs.Response {
	
	user, err := s.repo.GetEmail(body.Email)
	if err != nil {
		return libs.Respond("Email or password is incorrect", 401, true)
	}

	if libs.PasswordCheck(body.Password, user.Password) {
		return libs.Respond("Email or password is incorrect", 401, true)
	}

	if !user.IsVerified {
		return libs.Respond("You account is not verified", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)

	token, err := jwt.CreateToken()
	if err != nil {
		libs.Respond(err.Error(), 500, true)
	}

	return libs.Respond(tokenRes{Token: token}, 200, false)
}

//VERIFY EMAIL
func (s *Auth_Service) VerifyEmail(token string) *libs.Response {
	
	tokenExists := s.repo.TokenExists(token)
	if !tokenExists {
		return libs.Respond("Verification failed", 401, true)
	}

	user, err := s.repo.GetToken(token)
	if err != nil {
		return libs.Respond("User does not exist", 401, true)
	}

	if user.IsVerified {
		return libs.Respond("Your email has been registered", 401, true)
	}

	var data models.User

	data.IsVerified = true

	_, err = s.repo.UpdateUser(&data, user.UserID)
	if err != nil {
		return libs.Respond("User does not exists", 401, true)
	}

	res := map[string]string{"message": "Email has been verified"}

	return libs.Respond(res, 200, false)
}