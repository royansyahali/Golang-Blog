package repositories

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/payloads/request"
)

type PostHasCategoryRepository interface {
	InsertPostHasCategory(context.Context, *sql.Tx, *request.PostHasCategoryRequest) error
	UpdatePostHasCategory(context.Context, *sql.Tx, *request.PostHasCategoryRequest) error
	DeletePostHasCategory(context.Context, *sql.Tx, int) error
}
