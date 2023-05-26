package history

import (
	"github.com/gorilla/mux"
	"github.com/pius706975/backend/middlewares"
	"gorm.io/gorm"
)

func NewRouter(route *mux.Router, db *gorm.DB)  {
	router := route.PathPrefix("/history").Subrouter()

	repo := NewHistoryRepo(db)
	service := NewHistoryService(repo)
	ctrl := NewHistoryController(service)

	router.HandleFunc("/", middlewares.Handler(ctrl.GetAllHistories, middlewares.AuthMiddle("user"))).Methods("GET")
}