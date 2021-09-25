package repositories

import (
	"context"
	"database/sql"

	"blog/entities"
)

type AuthRepository interface {
	Login(context.Context, *sql.Tx, *entities.User) (entities.User, error)
}
