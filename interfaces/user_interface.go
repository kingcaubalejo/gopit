package interfaces

import (
	"go-api-jwt-v2/services/models"
)

type User interface {
	DisplayList(UUID int) (models.Users, error)
	Save(u models.Users) (error)
	Delete(u models.Users) (error)
	Update(u models.Users) (error)
}