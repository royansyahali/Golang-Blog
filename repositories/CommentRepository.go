package repositories

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/payloads/request"
)

type CommentRepository interface {
	InsertComment(context.Context, *sql.Tx, *request.CommentRequest) error
}
