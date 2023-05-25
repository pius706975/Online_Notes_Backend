package notes

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middlewares"
	"gorm.io/gorm"
)

func NewRouter(route *mux.Router, db *gorm.DB)  {
	router := route.PathPrefix("/note").Subrouter()

	repo := NewNoteRepo(db)
	service := NewNoteService(repo)
	ctrl := NewNoteController(service)

	router.HandleFunc("/add_note", middlewares.Handler(ctrl.AddNewNote, middlewares.AuthMiddle("user"))).Methods("POST")
}