package repositories

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/payloads/request"
)

type PostHasTagRepository interface {
	InsertPostHasTag(context.Context, *sql.Tx, *request.PostHasTagRequest) error
	UpdatePostHasTag(context.Context, *sql.Tx, *request.PostHasTagRequest) error
	DeletePostHasTag(context.Context, *sql.Tx, int) error
}
