package repositories

import (
	"context"
	"database/sql"

	"blog/entities"
	"blog/payloads/request"
	updaterequest "blog/payloads/request/updateRequest"
)

type UserRepository interface {
	InsertUser(context.Context, *sql.Tx, *request.UserRequest) error
	UpdateUser(context.Context, *sql.Tx, *updaterequest.UserUpdate, int) error
	DeleteUser(context.Context, *sql.Tx, int) error
	FindByIdUser(context.Context, *sql.Tx, int) (entities.User, error)
	GetAllUser(context.Context, *sql.Tx) ([]entities.User, error)
}
