package impl

import (
	"context"
	"database/sql"
	"errors"

	"blog/entities"
	"blog/payloads/request"
	"blog/repositories"
)

type CategoryRepositoryImp struct {
}

func NewCategoryRepository() repositories.CategoryRepository {
	return &CategoryRepositoryImp{}
}

func (c *CategoryRepositoryImp) InsertCategory(ctx context.Context, tx *sql.Tx, reqCategory *request.CategoryRequest) error {
	repoUser := NewUserRepository()
	user, err := repoUser.FindByIdUser(ctx, tx, reqCategory.UserId)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		query := ("INSERT INTO tb_categories (name) VALUES (?)")
		_, err = tx.ExecContext(ctx, query, reqCategory.Name)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("don't have permission to create this category")

}

func (c *CategoryRepositoryImp) FindByIdCategory(ctx context.Context, tx *sql.Tx, id int) (entities.Category, error) {
	category := entities.Category{}
	query := ("SELECT id, name FROM tb_categories WHERE id = ?")
	result, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return category, err
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(&category.Id,
			&category.Name,
		)
		if err != nil {
			return category, err
		}
		return category, err
	}

	return category, errors.New("category is not found in database")
}

func (p *CategoryRepositoryImp) GetAllCategory(ctx context.Context, tx *sql.Tx) ([]entities.Category, error) {
	var categories []entities.Category
	var category entities.Category
	query := ("SELECT id, name FROM tb_categories")
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&category.Id,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)

	}
	if len(categories) != 0 {
		return categories, nil
	}
	return categories, errors.New("category is empty")
}
