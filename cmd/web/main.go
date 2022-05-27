package main

import (
	"log"
	"net/http"

	"github.com/guibedin/poll/web/server"
	"github.com/guibedin/poll/web/service"
)

type storageType int

const (
	sql  storageType = 1
	file storageType = 2
)

func main() {
	server := server.New()

	// Setup service
	var svc service.Service
	svcType := sql

	switch svcType {
	case sql:
		svc = service.NewSqlService()
	case file:
		svc = service.NewFileService()
	}
	server.SetService(svc)

	log.Fatal(http.ListenAndServe(":8080", server))
}
