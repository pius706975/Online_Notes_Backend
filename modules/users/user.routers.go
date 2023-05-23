package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(route *mux.Router, db *gorm.DB)  {
	router := route.PathPrefix("/user").Subrouter()

	repo := NewUserRepo(db)
	service := NewUserService(repo)
	ctrl := NewUserController(service)

	router.HandleFunc("/register", ctrl.Register).Methods("POST")
}