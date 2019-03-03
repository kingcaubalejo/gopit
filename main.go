package main

import (
	"net/http"

	"go-api-jwt/routers"
	"go-api-jwt/settings"

	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5050", n)
}