package impl

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/repositories"
	"github.com/royansyahali/blog/services"
)

type TagServiceImpl struct {
	TagRepository repositories.TagRepository
	DB            *sql.DB
	Validasi      *validator.Validate
}

func NewTagService(p repositories.TagRepository, db *sql.DB, v *validator.Validate) services.TagService {
	return &TagServiceImpl{TagRepository: p, DB: db, Validasi: v}
}

func (t *TagServiceImpl) InsertTag(ctx context.Context, reqTag *request.TagRequest) error {
	err := t.Validasi.Struct(reqTag)
	if err != nil {
		return err
	}
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	err = t.TagRepository.InsertTag(ctx, tx, reqTag)
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

func (t *TagServiceImpl) FindByIdTag(ctx context.Context, id int) (entities.Tag, error) {
	var tag entities.Tag
	tx, err := t.DB.Begin()
	if err != nil {
		return tag, err
	}
	tag, err = t.TagRepository.FindByIdTag(ctx, tx, id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return tag, errRollback
		}
		return tag, err
	}
	err = tx.Commit()
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (t *TagServiceImpl) GetAllTag(ctx context.Context) ([]entities.Tag, error) {
	var tags []entities.Tag
	tx, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}
	tags, err = t.TagRepository.GetAllTag(ctx, tx)
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
	return tags, nil
}
