package routers

import(
	"go-api-jwt-v2/controllers"
	_"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth", controllers.AuthLogin).Methods("POST")
	
	// router.Handle(
	// 	"/refresh-token-auth",
	// 	negroni.New(
	// 		negroni.HandlerFunc(controllers.RefreshToken),
	// 	)).Methods("GET")

	// router.Handle(
	// 	"/logout",
	// 	negroni.New(
	// 		negroni.HandlerFunc(
	// 			// authentication.RequireTokenAuthentication,
	// 		),
	// 		negroni.HandlerFunc(controllers.Logout),
	// 	)).Methods("GET")
	return router
}