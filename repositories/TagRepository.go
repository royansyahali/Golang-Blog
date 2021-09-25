package repositories

import (
	"context"
	"database/sql"

	"blog/entities"
	"blog/payloads/request"
)

type TagRepository interface {
	InsertTag(context.Context, *sql.Tx, *request.TagRequest) error
	FindByIdTag(context.Context, *sql.Tx, int) (entities.Tag, error)
	GetAllTag(context.Context, *sql.Tx) ([]entities.Tag, error)
}
