package notes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(route *mux.Router, db *gorm.DB)  {
	// router := route.PathPrefix("/note").Subrouter()

	// repo := NewNoteRepo(db)
	// service := NewNoteService(repo)
	// ctrl := NewNoteController(service)
}