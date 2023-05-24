package users

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
	"github.com/pius706975/backend/middlewares"
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

// UPDATE
func (c *User_Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	ID := r.Context().Value(middlewares.UserID("user")).(string)

	var user models.User

	imageName := r.Context().Value("imageName").(string)
	user.Image = imageName

	err := schema.NewDecoder().Decode(&user, r.MultipartForm.Value)
	if err != nil {
		libs.Respond(err.Error(), 500, true).Send(w)
		return
	}

	c.svc.UpdateUser(&user, ID).Send(w)
}

// GET PROFILE
func (c *User_Controller) GetProfile(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	ID := r.Context().Value(middlewares.UserID("user")).(string)

	c.svc.GetProfile(ID).Send(w)
}