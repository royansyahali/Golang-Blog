package services

import (
	"context"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	updaterequest "github.com/royansyahali/blog/payloads/request/updateRequest"
)

type UserService interface {
	InsertUser(context.Context, *request.UserRequest) error
	UpdateUser(context.Context, *updaterequest.UserUpdate, int) error
	DeleteUser(context.Context, int) error
	FindByIdUser(context.Context, int) (entities.User, error)
	GetAllUser(context.Context) ([]entities.User, error)
}
