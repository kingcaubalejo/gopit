package controllers

import (
	"encoding/json"
	"net/http"

	"go-api-jwt-v2/services"
	"go-api-jwt-v2/services/models"
)

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.Users)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	
	authResult, statusCode, err := services.AuthUser(requestUser)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	
	var data = make(map[string]interface{})
	json.Unmarshal(authResult, &data)

	cookie := http.Cookie{Name: "Auth", Value: data["token"].(string),}
	http.SetCookie(w, &cookie)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(authResult)
}