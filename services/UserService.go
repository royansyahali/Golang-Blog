package services

import (
	"context"

	"blog/entities"
	"blog/payloads/request"
	updaterequest "blog/payloads/request/updateRequest"
)

type UserService interface {
	InsertUser(context.Context, *request.UserRequest) error
	UpdateUser(context.Context, *updaterequest.UserUpdate, int) error
	DeleteUser(context.Context, int) error
	FindByIdUser(context.Context, int) (entities.User, error)
	GetAllUser(context.Context) ([]entities.User, error)
}
