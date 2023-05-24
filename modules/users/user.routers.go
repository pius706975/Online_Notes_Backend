package users

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middlewares"
	"gorm.io/gorm"
)

func NewRouter(route *mux.Router, db *gorm.DB)  {
	router := route.PathPrefix("/user").Subrouter()

	repo := NewUserRepo(db)
	service := NewUserService(repo)
	ctrl := NewUserController(service)

	router.HandleFunc("/register", ctrl.Register).Methods("POST")

	router.HandleFunc("/profile/edit", middlewares.Handler(ctrl.UpdateUser, middlewares.AuthCloudUploadFile(), middlewares.AuthMiddle("admin", "user"))).Methods("PUT")
}