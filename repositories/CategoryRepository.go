package repositories

import (
	"context"
	"database/sql"

	"blog/entities"
	"blog/payloads/request"
)

type CategoryRepository interface {
	InsertCategory(context.Context, *sql.Tx, *request.CategoryRequest) error
	FindByIdCategory(context.Context, *sql.Tx, int) (entities.Category, error)
	GetAllCategory(context.Context, *sql.Tx) ([]entities.Category, error)
}
