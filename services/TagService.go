package services

import (
	"context"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
)

type TagService interface {
	InsertTag(context.Context, *request.TagRequest) error
	FindByIdTag(context.Context, int) (entities.Tag, error)
	GetAllTag(context.Context) ([]entities.Tag, error)
}
