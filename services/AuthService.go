package services

import (
	"context"

	"blog/entities"
)

type AuthService interface {
	Login(context.Context, *entities.User) error
}
