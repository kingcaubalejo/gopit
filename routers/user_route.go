package routers

import (
	"go-api-jwt-v2/controllers"
	"go-api-jwt-v2/lib/jwt"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.Handle("/user/list",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.UserDisplayList),
		)).Methods("GET")

	router.Handle("/user/data",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.UserDisplayListById),
		)).Methods("GET")

	router.Handle("/user/create",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.CreateUser),
		)).Methods("POST")

	router.Handle("/user/update",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.UpdateUser),
		)).Methods("PUT")

	router.Handle("/user/delete",
		negroni.New(
			negroni.HandlerFunc(jwt.JwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.DeleteUser),
		)).Methods("DELETE")

	return router
} 