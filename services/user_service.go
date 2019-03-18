package services

import (
	"fmt"
	"time"

	"go-api-jwt-v2/repository"
	"go-api-jwt-v2/services/models"
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

func CreateUser(user *models.Users) (int, []byte) {
	return 200, repository.CreateUser(user)
}

func UpdateUser(user *models.Users) (int, string) {
	repository.UpdateUser(user)
	return 200, "TEST"
}

func DeleteUser(user *models.Users) (int, string) {
	repository.DeleteUser(user)
	return 200, "TEST"
}

var ur repository.DbUserRepo

func DisplayUserList(uuid int) (int, models.Users) {
	list, err := ur.DisplayList(uuid)
	fmt.Println(err, "ERR")
	if err != nil {
		return 500, models.Users{}
	}
	return 200, list
}

type error interface {
	Error() string
}

type Errors struct {
	When time.Time
	What string
}

func (e *Errors) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func SaveUser(u models.Users) (int, error){
	err := ur.Save(u)
	if err != nil {
		return 500, &Errors{
			When: time.Now(),
			What: "it didn't work",
		}
	}
	return 200, nil
}