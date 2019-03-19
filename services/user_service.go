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

var ur repository.DbUserRepo

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

func DisplayListUser() (int, []models.Users) {
	list, err := ur.DisplayList()
	if err != nil {
		return 500, []models.Users{}
	}

	return 200, list
}

func DisplayListUserById(uuid int) (int, models.Users) {
	list, err := ur.DisplayListById(uuid)
	if err != nil {
		return 500, models.Users{}
	}

	return 200, list
}

func SaveUser(u models.Users) (int, error) {
	err := ur.Save(u)
	if err != nil {
		return 500, &Errors{
			When: time.Now(),
			What: "it didn't work",
		}
	}

	return 200, nil
}

func UpdateUser(u models.Users) (int, error) {
	err := ur.Update(u)
	if err != nil {
		return 500, &Errors{
			When: time.Now(),
			What: "it didn't work",
		}
	}

	return 200, nil
}

func DeleteUser(uuid int) (int, error) {
	err := ur.Delete(uuid)
	if err != nil {
		return 500, &Errors{
			When: time.Now(),
			What: "it didn't work",
		}
	}

	return 200, nil
}