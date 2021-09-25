package services

import (
	"context"

	"blog/entities"
	"blog/payloads/request"
)

type CategoryService interface {
	InsertCategory(context.Context, *request.CategoryRequest) error
	FindByIdCategory(context.Context, int) (entities.Category, error)
	GetAllCategory(context.Context) ([]entities.Category, error)
}
