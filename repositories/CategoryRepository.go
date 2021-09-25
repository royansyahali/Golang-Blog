package repositories

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
)

type CategoryRepository interface {
	InsertCategory(context.Context, *sql.Tx, *request.CategoryRequest) error
	FindByIdCategory(context.Context, *sql.Tx, int) (entities.Category, error)
	GetAllCategory(context.Context, *sql.Tx) ([]entities.Category, error)
}
