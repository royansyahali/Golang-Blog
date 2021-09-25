package impl

import (
	"context"
	"database/sql"

	"blog/payloads/request"
	"blog/repositories"
)

type PostHasTagRepositoryImp struct {
}

func NewPostHasTagRepository() repositories.PostHasTagRepository {
	return &PostHasTagRepositoryImp{}
}

func (t *PostHasTagRepositoryImp) InsertPostHasTag(ctx context.Context, tx *sql.Tx, reqPostHasTag *request.PostHasTagRequest) error {
	repoTag := NewTagRepository()
	tags, err := repoTag.GetAllTag(ctx, tx)
	if err != nil {
		return err
	}
	err = reqPostHasTag.Check(tags)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO tb_post_has_tag (post_id, tag_id) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range reqPostHasTag.Tag {
		_, err = stmt.ExecContext(ctx, reqPostHasTag.PostId, v)
		if err != nil {
			return err
		}
	}
	return nil

}

func (t *PostHasTagRepositoryImp) UpdatePostHasTag(ctx context.Context, tx *sql.Tx, reqPostHasTag *request.PostHasTagRequest) error {
	err := t.DeletePostHasTag(ctx, tx, reqPostHasTag.PostId)
	if err != nil {
		return err
	}
	err = t.InsertPostHasTag(ctx, tx, reqPostHasTag)
	if err != nil {
		return err
	}
	return nil

}

func (t *PostHasTagRepositoryImp) DeletePostHasTag(ctx context.Context, tx *sql.Tx, id int) error {
	query := ("DELETE FROM tb_post_has_tag WHERE post_id = ?")
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil

}
