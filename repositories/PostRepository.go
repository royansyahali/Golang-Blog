package repositories

import (
	"context"
	"database/sql"

	"blog/payloads/request"
	updaterequest "blog/payloads/request/updateRequest"
	"blog/payloads/response"
)

type PostRepository interface {
	InsertPost(context.Context, *sql.Tx, *request.PostRequest) error
	UpdatePost(context.Context, *sql.Tx, *updaterequest.PostUpdate, int) error
	DeletePost(context.Context, *sql.Tx, int, int) error
	FindByIdPost(context.Context, *sql.Tx, int) (response.PostResponse, error)
	GetAllPost(context.Context, *sql.Tx) ([]response.PostResponse, error)
}
