package impl

import (
	"context"
	"database/sql"

	"blog/entities"
	"blog/payloads/request"
	"blog/repositories"
	"blog/services"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	Validasi           *validator.Validate
}

func NewCategoryService(p repositories.CategoryRepository, db *sql.DB, v *validator.Validate) services.CategoryService {
	return &CategoryServiceImpl{CategoryRepository: p, DB: db, Validasi: v}
}

func (c *CategoryServiceImpl) InsertCategory(ctx context.Context, reqCategory *request.CategoryRequest) error {
	err := c.Validasi.Struct(reqCategory)
	if err != nil {
		return err
	}
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	err = c.CategoryRepository.InsertCategory(ctx, tx, reqCategory)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryServiceImpl) FindByIdCategory(ctx context.Context, id int) (entities.Category, error) {
	var Category entities.Category
	tx, err := c.DB.Begin()
	if err != nil {
		return Category, err
	}
	Category, err = c.CategoryRepository.FindByIdCategory(ctx, tx, id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return Category, errRollback
		}
		return Category, err
	}
	err = tx.Commit()
	if err != nil {
		return Category, err
	}
	return Category, nil
}

func (c *CategoryServiceImpl) GetAllCategory(ctx context.Context) ([]entities.Category, error) {
	var categories []entities.Category
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}
	categories, err = c.CategoryRepository.GetAllCategory(ctx, tx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return nil, errRollback
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
