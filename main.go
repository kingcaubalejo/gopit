package main

import (
	"net/http"
	"fmt"

	"go-api-jwt-v2/routers"
	"go-api-jwt-v2/settings"
	"go-api-jwt-v2/repository"

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