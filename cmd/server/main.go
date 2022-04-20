package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.Handler()
	log.Fatal(http.ListenAndServe(":8080", router))
}
