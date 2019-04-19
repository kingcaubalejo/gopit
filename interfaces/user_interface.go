package interfaces

import (
	"go-api-jwt-v2/services/models"
)

type User interface {
	DisplayList() (models.Users, error)
	DisplayListById(uuid int) (models.Users, error)
	Save(u models.Users) (error)
	Delete(uuid int) (error)
	Update(u models.Users) (error)
	DeleteMultiple(u []models.Users) (error)
}