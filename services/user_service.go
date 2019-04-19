package services

import (
	"fmt"
	_"time"
	"encoding/json"
	_"errors"
	
	"go-api-jwt-v2/repository"
	"go-api-jwt-v2/services/models"
)

var ur repository.DbUserRepo

func DisplayListUser() ([]byte, int, error) {
	list, err := ur.DisplayList()
	if err != nil {
		return nil, 500, err
	}

	resultParsed, errList := json.Marshal(list)
	if errList != nil {
		return nil, 500, errList
	}
	return resultParsed, 200, nil
}

func DisplayListUserById(uuid int) ([]byte, int, error) {
	list, err := ur.DisplayListById(uuid)
	if err != nil {
		return []byte(""), 500, err
	}

	resultParsed, errList := json.Marshal(list)
	if errList != nil {
		return nil, 500, errList
	}
	return resultParsed, 200, nil
}

func SaveUser(u models.Users) ([]byte, int, error) {
	err := ur.Save(u)
	if err != nil {
		return []byte(""), 500, err
	}

	resultParsed, errList := json.Marshal(map[string]interface{}{
		"data": "User is successfully created.",
	})
	if errList != nil {
		return nil, 500, errList
	}

	return resultParsed, 200, nil
}

func UpdateUser(u models.Users) ([]byte, int, error) {
	err := ur.Update(u)
	if err != nil {
		return []byte(""), 500, err
	}

	resultParsed, errList := json.Marshal(map[string]interface{}{
		"data": "User is successfully updated.",
	})
	if errList != nil {
		return nil, 500, errList
	}
	
	return resultParsed, 200, nil
}

func DeleteUser(uuid int) ([]byte, int, error) {
	err := ur.Delete(uuid)
	if err != nil {
		return []byte(""), 500, err
	}

	resultParsed, errList := json.Marshal(map[string]interface{}{
		"data": "User is successfully deleted.",
	})
	if errList != nil {
		return nil, 500, errList
	}
	
	return resultParsed, 200, nil
}

func DeleteMultipleUser(uuid []int) ([]byte, int, error) {
	fmt.Println(uuid, "services multiple delete")
	err := ur.DeleteMultiple(uuid)
	if err != nil {
		return []byte(""), 500, err
	}

	resultParsed, errList := json.Marshal(map[string]interface{}{
		"data": "User(s) is successfully deleted.",
	})
	if errList != nil {
		return nil, 500, errList
	}
	
	return resultParsed, 200, nil
}