package services

import (
	"encoding/json"
	_"net/http"
	_"fmt"
	_ "strconv"

	_"go-scaffolding/api/parameters"
	"go-api-jwt-v2/lib/jwt"
	_"go-scaffolding/core/authentication"
	"go-api-jwt-v2/services/models"
	"go-api-jwt-v2/repository"
)

var auth repository.DbAuthRepo

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

func AuthUser(u *models.Users) ([]byte, int, error) {
	authResult, err := auth.AuthenticateUser(u)
	if err != nil {
		return []byte(""), 500, err
	}

	result, _ := json.Marshal(authResult)
	return result, 200, nil
}