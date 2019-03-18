package interfaces

import (
	"go-api-jwt-v2/services/models"
)

type UserSaver interface {
	DisplayList(UUID int) (models.Users, error)
}