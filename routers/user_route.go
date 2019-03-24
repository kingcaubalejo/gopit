package routers

import (
	"go-api-jwt-v2/controllers"
	_"go-scaffolding/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-api-jwt-v2/lib/jwt"
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

	router.HandleFunc("/user/list",
		controllers.UserDisplayList,
	).Methods("GET")

	router.HandleFunc("/user/data",
		controllers.UserDisplayListById,
	).Methods("GET")

	router.HandleFunc("/user/create",
		controllers.CreateUser,
	).Methods("POST")

	router.HandleFunc("/user/update",
		controllers.UpdateUser,
	).Methods("PUT")

	router.HandleFunc("/user/delete",
		controllers.DeleteUser,
	).Methods("DELETE")

	return router
} 