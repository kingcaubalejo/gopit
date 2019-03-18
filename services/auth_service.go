package services

import (
	_"encoding/json"
	_"net/http"
	_"fmt"
	_ "strconv"

	_"go-scaffolding/api/parameters"
	"go-api-jwt-v2/lib/jwt"
	_"go-scaffolding/core/authentication"
	"go-api-jwt-v2/services/models"
)

func Login(requestUser *models.Users) (int, string){
	
	cipher := jwt.Crackdependmaker(string(requestUser.UUID))
	responseStatus, token := jwt.SignToken(cipher, "access_token")
	return responseStatus, token
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