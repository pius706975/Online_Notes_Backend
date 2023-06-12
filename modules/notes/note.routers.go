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
	
	router.HandleFunc("/", middlewares.Handler(ctrl.GetAllNotes, middlewares.AuthMiddle("user"))).Methods("GET")
	router.HandleFunc("/{id}", middlewares.Handler(ctrl.GetByID, middlewares.AuthMiddle("user"))).Methods("GET")
	router.HandleFunc("/search/{query}", middlewares.Handler(ctrl.SearchNote, middlewares.AuthMiddle("user"))).Methods("GET")

	router.HandleFunc("/add_note", middlewares.Handler(ctrl.AddNewNote, middlewares.AuthMiddle("user"))).Methods("POST")

	router.HandleFunc("/edit/{id}", middlewares.Handler(ctrl.UpdateNote, middlewares.AuthMiddle("user"))).Methods("PUT")

	router.HandleFunc("/delete/{id}", middlewares.Handler(ctrl.DeleteNote, middlewares.AuthMiddle("user"))).Methods("DELETE")
}