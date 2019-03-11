package main

import (
	"net/http"
	"fmt"

	"go-api-jwt/routers"
	"go-api-jwt/settings"
	"go-api-jwt/repository"

	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	repository.InitDatabase()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	
	connection := fmt.Sprintf(":%s", settings.Get().ServerPort)
	http.ListenAndServe(connection, n)
}