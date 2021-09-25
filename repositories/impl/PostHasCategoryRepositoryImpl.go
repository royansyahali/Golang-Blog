package impl

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/repositories"
)

type PostHasCategoryRepositoryImp struct {
}

func NewPostHasCategoryRepository() repositories.PostHasCategoryRepository {
	return &PostHasCategoryRepositoryImp{}
}

func (t *PostHasCategoryRepositoryImp) InsertPostHasCategory(ctx context.Context, tx *sql.Tx, reqPostHasCategory *request.PostHasCategoryRequest) error {
	repoCategory := NewCategoryRepository()
	categories, err := repoCategory.GetAllCategory(ctx, tx)
	if err != nil {
		return err
	}
	err = reqPostHasCategory.Check(categories)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO tb_post_has_category (post_id,category_id) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range reqPostHasCategory.Category {
		_, err = stmt.ExecContext(ctx, reqPostHasCategory.PostId, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *PostHasCategoryRepositoryImp) UpdatePostHasCategory(ctx context.Context, tx *sql.Tx, reqPostHasCategory *request.PostHasCategoryRequest) error {
	err := t.DeletePostHasCategory(ctx, tx, reqPostHasCategory.PostId)
	if err != nil {
		return err
	}
	err = t.InsertPostHasCategory(ctx, tx, reqPostHasCategory)
	if err != nil {
		return err
	}
	return nil
}

func (t *PostHasCategoryRepositoryImp) DeletePostHasCategory(ctx context.Context, tx *sql.Tx, id int) error {
	query := ("DELETE FROM tb_post_has_category WHERE post_id = ?")
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil

}
