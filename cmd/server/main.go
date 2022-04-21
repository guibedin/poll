package main

import (
	"log"
	"net/http"

	"github.com/guibedin/voting/internal/create"
	http2 "github.com/guibedin/voting/internal/http"
	"github.com/guibedin/voting/internal/read"
	"github.com/guibedin/voting/internal/vote"
)

func main() {

	var creator create.Service
	var reader read.Service
	var voter vote.Service

	router := http2.Handler()
	log.Fatal(http.ListenAndServe(":8080", router))
}
