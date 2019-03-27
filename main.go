package main

import (
	"net/http"
	"fmt"

	"go-api-jwt-v2/routers"
	"go-api-jwt-v2/settings"
	"go-api-jwt-v2/repository"

	"github.com/gorilla/handlers"
	"github.com/codegangsta/negroni"
)

func main() {

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})

	settings.Init()
	repository.InitDatabase()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	
	connection := fmt.Sprintf(":%s", settings.Get().ServerPort)
	http.ListenAndServe(connection, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(n))
}