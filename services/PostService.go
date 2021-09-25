package services

import (
	"context"

	"blog/payloads/request"
	updaterequest "blog/payloads/request/updateRequest"
	"blog/payloads/response"
)

type PostService interface {
	InsertPost(context.Context, *request.PostRequest) error
	UpdatePost(context.Context, *updaterequest.PostUpdate, int) error
	DeletePost(context.Context, int, int) error
	FindByIdPost(context.Context, int) (response.PostResponse, error)
	GetAllPost(context.Context) ([]response.PostResponse, error)
}
