package impl

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/repositories"
)

type CommentRepositoryImp struct {
}

func NewCommentRepository() repositories.CommentRepository {
	return &CommentRepositoryImp{}
}

func (c *CommentRepositoryImp) InsertComment(ctx context.Context, tx *sql.Tx, reqComment *request.CommentRequest) error {
	repoUser := NewUserRepository()
	_, err := repoUser.FindByIdUser(ctx, tx, reqComment.UserId)
	if err != nil {
		return err
	}
	repoPost := NewPostRepository()
	_, err = repoPost.FindByIdPost(ctx, tx, reqComment.PostId)
	if err != nil {
		return err
	}
	query := ("INSERT INTO tb_comments (post_id, user_id, content, created_at, updated_at) VALUES (?,?,?,?,?)")
	_, err = tx.ExecContext(ctx, query,
		reqComment.PostId,
		reqComment.UserId,
		reqComment.Content,
		reqComment.CreateAt,
		reqComment.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil

}
