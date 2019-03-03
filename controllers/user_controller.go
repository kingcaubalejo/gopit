package controllers

import (
	"net/http"
	_"fmt"

	"go-api-jwt/services"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World"))
}

func TestLangController(w http.ResponseWriter, r *http.Request){
	services.TestLang()
	w.Write([]byte("DB binding..."))
}

func SelectData(w http.ResponseWriter, r *http.Request) {
	services.SelectData()
	w.Write([]byte("GG WELL PLAYED"))
}

func SelectWhereData(w http.ResponseWriter, r *http.Request) {
	services.SelectWhereData()
	w.Write([]byte("GG WELL PLAYED"))
}