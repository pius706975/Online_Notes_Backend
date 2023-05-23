package users

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
)

type User_Controller struct {
	svc *User_Service
}

func NewUserController(svc *User_Service) User_Controller {
	return User_Controller{svc}
}

//REGISTER
func (c *User_Controller) Register(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.Respond(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		libs.Respond(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.Register(&user).Send(w)
}