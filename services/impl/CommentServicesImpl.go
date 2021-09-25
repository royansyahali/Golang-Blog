package impl

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/repositories"
	"github.com/royansyahali/blog/services"
)

type CommentServiceImpl struct {
	CommentRepository repositories.CommentRepository
	DB                *sql.DB
	Validasi          *validator.Validate
}

func NewCommentService(p repositories.CommentRepository, db *sql.DB, v *validator.Validate) services.CommentService {
	return &CommentServiceImpl{CommentRepository: p, DB: db, Validasi: v}
}

func (c *CommentServiceImpl) InsertComment(ctx context.Context, reqComment *request.CommentRequest) error {
	err := c.Validasi.Struct(reqComment)
	if err != nil {
		return err
	}
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	err = c.CommentRepository.InsertComment(ctx, tx, reqComment)
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
