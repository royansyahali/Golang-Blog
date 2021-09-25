package repositories

import (
	"context"
	"database/sql"

	"blog/payloads/request"
)

type CommentRepository interface {
	InsertComment(context.Context, *sql.Tx, *request.CommentRequest) error
}
