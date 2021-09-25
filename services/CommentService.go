package services

import (
	"context"

	"blog/payloads/request"
)

type CommentService interface {
	InsertComment(context.Context, *request.CommentRequest) error
}
