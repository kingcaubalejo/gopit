package interfaces

import (
	"go-api-jwt-v2/services/models"
)

type Auth interface {
	AuthenticateUser() (models.Users, error)
}