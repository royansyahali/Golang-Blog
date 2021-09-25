package services

import (
	"context"

	"github.com/royansyahali/blog/entities"
)

type AuthService interface {
	Login(context.Context, *entities.User) error
}
