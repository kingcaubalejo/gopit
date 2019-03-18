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

	router.HandleFunc("/test/select",
		controllers.SelectData,
	).Methods("GET")

	router.HandleFunc("/test/select-where",
		controllers.SelectWhereData,
	).Methods("GET")

	router.HandleFunc("/test/create-user",
		controllers.CreateUser,
	).Methods("POST")

	router.HandleFunc("/test/update-user",
		controllers.UpdateUser,
	).Methods("POST")

	router.HandleFunc("/test/delete-user",
		controllers.DeleteUser,
	).Methods("POST")

	return router
} 