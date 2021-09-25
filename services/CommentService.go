package services

import (
	"context"

	"github.com/royansyahali/blog/payloads/request"
)

type CommentService interface {
	InsertComment(context.Context, *request.CommentRequest) error
}
