package services

import (
	"fmt"

	"go-api-jwt/repository"
)

func TestLang() (int, string){
	repository.DbConn()
	// cipher := jwt.Crackdependmaker(string(requestUser.UUID))
	// responseStatus, token := jwt.SignToken(cipher, "access_token")
	return 200, "ILOVE YOU"
	// authBackend := authentication.InitJWTAuthenticationBackend()

	// if authBackend.Authenticate(requestUser) {
		// token, err := authBackend.GenerateToken(requestUser.UUID)
		// token, err := authBackend.GenerateToken(requestUser.UUID)
		// fmt.Println(token, "HEY TOKEN")
		// if err != nil {
		// 	return http.StatusInternalServerError, []byte("")
		// } else {
		// 	response, _ := json.Marshal(parameters.TokenAuthentication{token})
		// 	return http.StatusOK, response
		// }
	// }
	// return http.StatusUnauthorized, []byte("")
}

func SelectData() (int, string) {
	fmt.Println("Select Service")
	repository.Select()
	return 200, "TEST"
}