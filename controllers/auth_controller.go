package controllers

import (
	"go-api-jwt/services"
	"go-api-jwt/services/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.Users)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	
	responseStatus, token := services.Login(requestUser)
	responseTokenAuth, _ := json.Marshal(map[string]interface{}{
		"token"	: token,
	})

    cookie := http.Cookie{Name: "Auth", Value: token}
    http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseTokenAuth)

}