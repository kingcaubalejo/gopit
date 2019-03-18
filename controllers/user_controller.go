package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"
	_"fmt"

	"go-api-jwt-v2/services"
	"go-api-jwt-v2/services/models"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World"))
}

func TestLangController(w http.ResponseWriter, r *http.Request){
	services.TestLang()
	w.Write([]byte("DB binding..."))
}

func SelectData(w http.ResponseWriter, r *http.Request) {
	statusCode, resultData := services.SelectData()
	resultDataParsed, _ := json.Marshal(resultData) 
	if statusCode == 200 {
		w.Write([]byte(resultDataParsed))	
	}
	
}

func SelectWhereData(w http.ResponseWriter, r *http.Request) {
	UUId, _ := strconv.Atoi(r.URL.Query().Get("uuid"))
	statusCode, resultData := services.DisplayUserList(UUId)
	resultDataParsed, _ := json.Marshal(resultData) 
	if statusCode == 200 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(resultDataParsed))	
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestCreateUser := models.Users{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestCreateUser)

	statusCode, respondResult := services.SaveUser(requestCreateUser)
	if statusCode == 200 && respondResult == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte("User is created successfully."))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestUpdateUser := new(models.Users)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUpdateUser)

	services.UpdateUser(requestUpdateUser)
	w.Write([]byte("GG WELL PLAYED"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	requestDeleteUser := new(models.Users)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestDeleteUser)

	services.DeleteUser(requestDeleteUser)
	w.Write([]byte("GG WELL PLAYED"))
}