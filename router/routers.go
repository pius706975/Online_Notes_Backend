package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pius706975/backend/database"
	"github.com/pius706975/backend/modules/auth"
	"github.com/pius706975/backend/modules/users"
)

func RouterApp() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := database.NewDB() // change _ to `db` when you want to use it

	if err != nil {
		return nil, err
	}

	subRouter := mainRoute.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/", homeHandler).Methods("GET")

	users.NewRouter(subRouter, db)
	auth.NewRouter(subRouter, db)

	return mainRoute, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Online Notes backend!"))
}
