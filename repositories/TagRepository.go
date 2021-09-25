package repositories

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
)

type TagRepository interface {
	InsertTag(context.Context, *sql.Tx, *request.TagRequest) error
	FindByIdTag(context.Context, *sql.Tx, int) (entities.Tag, error)
	GetAllTag(context.Context, *sql.Tx) ([]entities.Tag, error)
}
