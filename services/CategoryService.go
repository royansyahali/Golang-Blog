package services

import (
	"context"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
)

type CategoryService interface {
	InsertCategory(context.Context, *request.CategoryRequest) error
	FindByIdCategory(context.Context, int) (entities.Category, error)
	GetAllCategory(context.Context) ([]entities.Category, error)
}
