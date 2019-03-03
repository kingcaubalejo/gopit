package routers

import (
	"go-api-jwt/controllers"
	_"go-scaffolding/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-api-jwt/lib/jwt"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		"/test/hello",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.Wrap(jwt.MyHandler),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	router.HandleFunc("/test/db",
			controllers.TestLangController,
		).Methods("GET")

	router.HandleFunc("/test/select",
		controllers.SelectData,
	).Methods("GET")

	return router
} 