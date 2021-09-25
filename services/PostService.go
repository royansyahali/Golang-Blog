package services

import (
	"context"

	"github.com/royansyahali/blog/payloads/request"
	updaterequest "github.com/royansyahali/blog/payloads/request/updateRequest"
	"github.com/royansyahali/blog/payloads/response"
)

type PostService interface {
	InsertPost(context.Context, *request.PostRequest) error
	UpdatePost(context.Context, *updaterequest.PostUpdate, int) error
	DeletePost(context.Context, int, int) error
	FindByIdPost(context.Context, int) (response.PostResponse, error)
	GetAllPost(context.Context) ([]response.PostResponse, error)
}
