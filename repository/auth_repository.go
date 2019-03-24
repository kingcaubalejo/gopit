package repository

import (
	"database/sql"
   _ "fmt"
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
    if u.Username != "" {
        return nil, errors.New("Please insert username.")
    }

    if u.Password != "" {
        return nil, errors.New("Please insert password.")
    }

	authUser := models.Users{}
    errUsers := database.QueryRow("SELECT uuid, username FROM users WHERE username=? AND password=?", u.Username, u.Password).Scan(&authUser.UUID, &authUser.Username)

	cipher := lib.EncryptPlainText(string(authUser.UUID))
    responseStatus, token := jwt.SignToken(cipher, "access_token")
    
    if responseStatus != 200 {
        return nil, errors.New("Token is not created.")
    }

    if errUsers != nil {
        if errUsers != sql.ErrNoRows {
            return nil, errUsers
        }
        return nil, errors.New("Zero rows found.")
    }
    
    resultAuthenticate := map[string]interface{}{
        "uuid"          : authUser.UUID,
        "user_name"     : authUser.Username,
        "token"         : token,
    }

    return resultAuthenticate, nil
}