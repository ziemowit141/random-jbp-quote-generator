package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ziemowit141/random-jbp-quote-generator/src/server/handlers"
)

type Server struct {}

func (s Server) Start() {
	router := mux.NewRouter()

	router.Handle("/", handlers.NoMotivationHandler{})
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on 127.0.0.1:8000")
	log.Fatal(srv.ListenAndServe())
}