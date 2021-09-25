package impl

import (
	"context"
	"database/sql"

	"blog/payloads/request"
	updaterequest "blog/payloads/request/updateRequest"
	"blog/payloads/response"
	"blog/repositories"
	"blog/services"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repositories.PostRepository
	DB             *sql.DB
	Validasi       *validator.Validate
}

func NewPostService(p repositories.PostRepository, db *sql.DB, v *validator.Validate) services.PostService {
	return &PostServiceImpl{PostRepository: p, DB: db, Validasi: v}
}

func (p *PostServiceImpl) InsertPost(ctx context.Context, reqPost *request.PostRequest) error {
	err := p.Validasi.Struct(reqPost)
	if err != nil {
		return err
	}
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	err = p.PostRepository.InsertPost(ctx, tx, reqPost)
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

func (p *PostServiceImpl) UpdatePost(ctx context.Context, reqPost *updaterequest.PostUpdate, id int) error {
	err := p.Validasi.Struct(reqPost)
	if err != nil {
		return err
	}
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	err = p.PostRepository.UpdatePost(ctx, tx, reqPost, id)
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

func (p *PostServiceImpl) DeletePost(ctx context.Context, id_user, id int) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	err = p.PostRepository.DeletePost(ctx, tx, id_user, id)
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

func (p *PostServiceImpl) FindByIdPost(ctx context.Context, id int) (response.PostResponse, error) {
	var post response.PostResponse
	tx, err := p.DB.Begin()
	if err != nil {
		return post, err
	}
	post, err = p.PostRepository.FindByIdPost(ctx, tx, id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return post, errRollback
		}
		return post, err
	}
	err = tx.Commit()
	if err != nil {
		return post, err
	}
	return post, nil
}

func (p *PostServiceImpl) GetAllPost(ctx context.Context) ([]response.PostResponse, error) {
	var Posts []response.PostResponse
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	Posts, err = p.PostRepository.GetAllPost(ctx, tx)
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
	return Posts, nil
}
