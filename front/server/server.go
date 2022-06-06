package server

import (
	"log"
	"net/http"
)

func New() {

	http.Handle("/", http.FileServer(http.Dir("../../front/static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
