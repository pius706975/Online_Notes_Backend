package notes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/libs"
)

type Note_Controller struct {
	svc *Note_Service
}

func NewNoteController(svc *Note_Service) Note_Controller {
	return Note_Controller{svc}
}

// ADD NOTE
func (c *Note_Controller) AddNewNote(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		fmt.Println(err)
		libs.Respond(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&note)
	if err != nil {
		libs.Respond(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.AddNewNote(&note).Send(w)
}

// DELETE NOTE
func (c *Note_Controller) DeleteNote(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	idStr := mux.Vars(r)
	id, ok := idStr["id"]
	if !ok {
		libs.Respond("Get the ID", 400, true).Send(w)
		return
	}

	c.svc.DeleteNote(id).Send(w)
}

// GET ALL NOTES
func (c *Note_Controller) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllNotes().Send(w)
}

// SEARCH NOTE
func (c *Note_Controller) SearchNote(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	query := mux.Vars(r)["query"]
	query = strings.ToLower(query)

	c.svc.SearchNote(query).Send(w)
}