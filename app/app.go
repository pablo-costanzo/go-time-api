package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/time", getTime)

	//Start the server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
