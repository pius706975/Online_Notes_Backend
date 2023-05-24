package libs

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pius706975/backend/router"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "for running API server",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},

		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {

	mainRoute, err := router.RouterApp()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:3030"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "0.0.0.0:" + PORT //change this address to 0.0.0.0 if it's used in docker
	}

	cors := corsHandler()

	serve := &http.Server{
		Addr:         address,
		WriteTimeout: time.Minute * 2,
		ReadTimeout:  time.Minute * 2,
		IdleTimeout:  time.Minute,
		Handler:      cors.Handler(mainRoute),
	}

	log.Println("App is running on PORT 3030")

	return serve.ListenAndServe()
}
