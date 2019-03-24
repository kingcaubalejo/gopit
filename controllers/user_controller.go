package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"go-api-jwt-v2/services"
	"go-api-jwt-v2/services/models"
)

func UserDisplayList(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	resultData, statusCode, err := services.DisplayListUser()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resultData)
}

func UserDisplayListById(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	UUId, _ := strconv.Atoi(r.URL.Query().Get("uuid"))
	resultData, statusCode, err := services.DisplayListUserById(UUId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resultData)	
}

func CreateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestCreateUser := models.Users{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestCreateUser)

	resultData, statusCode, err := services.SaveUser(requestCreateUser)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resultData)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUpdateUser := models.Users{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUpdateUser)

	resultData, statusCode, err := services.UpdateUser(requestUpdateUser)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resultData)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	UUId, _ := strconv.Atoi(r.URL.Query().Get("uuid"))

	resultData, statusCode, err := services.DeleteUser(UUId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resultData)
}