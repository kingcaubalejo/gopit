package services

import (
	_ "fmt"

	"go-api-jwt/repository"
	"go-api-jwt/services/models"
)

func TestLang() (int, string){
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

func SelectData() (int, map[string]interface{}) {
	return 200, repository.SelectAll()
}

func SelectWhereData(uuid int) (int, map[string]interface{}) {
	return 200, repository.SelectWhere(uuid)
}

func CreateUser(user *models.Users) (int, string) {
	repository.CreateUser(user)
	return 200, "TEST"
}

func UpdateUser(user *models.Users) (int, string) {
	repository.UpdateUser(user)
	return 200, "TEST"
}

func DeleteUser(user *models.Users) (int, string) {
	repository.DeleteUser(user)
	return 200, "TEST"
}