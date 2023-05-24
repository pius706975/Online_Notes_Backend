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

	router.HandleFunc("/profile", middlewares.Handler(ctrl.GetProfile, middlewares.AuthMiddle("admin", "user"))).Methods("GET")
	router.HandleFunc("/all_users", middlewares.Handler(ctrl.GetAllUsers, middlewares.AuthMiddle("admin"))).Methods("GET")

	router.HandleFunc("/register", ctrl.Register).Methods("POST")

	router.HandleFunc("/profile/edit", middlewares.Handler(ctrl.UpdateUser, middlewares.AuthCloudUploadFile(), middlewares.AuthMiddle("admin", "user"))).Methods("PUT")

	router.HandleFunc("/profile/delete", middlewares.Handler(ctrl.DeleteUser, middlewares.AuthMiddle("admin", "user"))).Methods("DELETE")
}