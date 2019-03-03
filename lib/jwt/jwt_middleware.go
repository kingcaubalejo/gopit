
package jwt

import (
	jwtAuthenticate "github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"

	"net/http"
	"fmt"
	_ "strconv"
)

var MyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user");
	fmt.Fprintf(w, "This is an authenticated request")
	fmt.Fprintf(w, "Claim content:\n")
	fmt.Println(user, "USER")
	// fmt.Println(len(user.(*jwtAuthenticate.Token).Claims), "User MyHandler")
	// for k, v := range user.(*jwt.Token).Claims {
	// 	fmt.Fprintf(w, "%s :\t%#v\n", k, v)
	// }
	// cookie, _ := r.Cookie("Auth")
	// fmt.Println(cookie, "cookie", r)
	// //ErrorHandler errorHandler
	// //jData, _ := json.Marshal("Required authorization token not found")
	// if cookie.Value == "none" {
	// 	// w.WriteHeader(401)
	//     // w.Write(jData)
	// 	jwtmiddleware.OnError(w,r,"Required authorization token not found")
	// }
	// return
})
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
      ValidationKeyGetter : func(token *jwtAuthenticate.Token) (interface{}, error) {
		authBackend := InitJWTAuthenticationBackend()
		return authBackend.PublicKey, nil
      },
      SigningMethod: jwtAuthenticate.SigningMethodRS512,
})
