package services

import (
	"encoding/json"

	"go-api-jwt-v2/services/models"
	"go-api-jwt-v2/repository"
)

var auth repository.DbAuthRepo

func AuthUser(u *models.Users) ([]byte, int, error) {
	authResult, err := auth.AuthenticateUser(u)
	if err != nil {
		return []byte(""), 500, err
	}

	result, _ := json.Marshal(authResult)
	return result, 200, nil
}