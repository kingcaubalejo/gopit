package repository

import (
	"database/sql"
    _"fmt"
    "errors"
   
    "go-api-jwt-v2/services/models"
    "go-api-jwt-v2/interfaces"
    "go-api-jwt-v2/lib/jwt"
    "go-api-jwt-v2/lib"
)


type DbAuthRepo struct {
   Repository interfaces.Auth
}

func (auth *DbAuthRepo) AuthenticateUser(u *models.Users) (map[string]interface{}, error) {
    if u.Username == "" || u.Password == "" {
        return nil, errors.New("Invalid username/password.")
    }

    password := lib.EncryptPlainText(u.Username + u.Password)
	authUser := models.Users{}
    errUsers := database.QueryRow("SELECT uuid, username FROM users WHERE username=? AND password=?", u.Username, password).Scan(&authUser.UUID, &authUser.Username)

    if errUsers != nil {
        if errUsers != sql.ErrNoRows {
            return nil, errUsers
        }
        return nil, errors.New("Invalid username/password.")
    }

    cipher := lib.EncryptPlainText(string(authUser.UUID))
    responseStatus, token := jwt.SignToken(cipher, "access_token")
    
    if responseStatus != 200 {
        return nil, errors.New("Token is not created.")
    }

    resultAuthenticate := map[string]interface{}{
        "uuid"          : authUser.UUID,
        "user_name"     : authUser.Username,
        "token"         : token,
    }
    return resultAuthenticate, nil
}