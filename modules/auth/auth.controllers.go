package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
)

type Auth_Controller struct {
	svc Auth_Service
}

func NewAuthController(svc Auth_Service) *Auth_Controller {
	return &Auth_Controller{svc}
}

//LOGIN
func (c *Auth_Controller) Login(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.Respond(err.Error(), 500, true).Send(w)
		return
	}

	c.svc.Login(&user).Send(w)
}

// VERIFIY EMAIL
func (c *Auth_Controller) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	mux := mux.Vars(r)

	token, ok := mux["token"]
	if !ok {
		libs.Respond("Token not found", 404, true).Send(w)
		return
	}

	c.svc.VerifyEmail(token).Send(w)
}