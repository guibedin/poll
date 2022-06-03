package main

import (
	"log"
	"net/http"

	"github.com/guibedin/poll/web/repository"
	"github.com/guibedin/poll/web/server"
	"github.com/guibedin/poll/web/service"
)

func main() {
	server := server.New()

	// Set repository type
	repoType := repository.Sql

	// Setup service
	svc := service.New(repository.New(repoType))

	// Setup server
	server.SetService(svc)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", server))
}
