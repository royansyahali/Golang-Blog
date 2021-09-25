package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/repositories"
)

type TagRepositoryImp struct {
}

func NewTagRepository() repositories.TagRepository {
	return &TagRepositoryImp{}
}

func (t *TagRepositoryImp) InsertTag(ctx context.Context, tx *sql.Tx, reqTag *request.TagRequest) error {
	repoUser := NewUserRepository()
	user, err := repoUser.FindByIdUser(ctx, tx, reqTag.UserId)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		query := ("INSERT INTO tb_tags (name) VAlUES (?)")
		result, err := tx.QueryContext(ctx, query, reqTag.Name)
		if err != nil {
			return err
		}
		defer result.Close()

		return nil
	}
	return errors.New("don't have permission to create this tag")

}

func (t *TagRepositoryImp) FindByIdTag(ctx context.Context, tx *sql.Tx, id int) (entities.Tag, error) {
	tag := entities.Tag{}
	query := ("SELECT id, name FROM tb_tags WHERE id = ?")
	result, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return tag, err
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(&tag.Id,
			&tag.Name,
		)
		if err != nil {
			return tag, err
		}
		return tag, err
	}

	return tag, errors.New("tag is not found in database")
}

func (p *TagRepositoryImp) GetAllTag(ctx context.Context, tx *sql.Tx) ([]entities.Tag, error) {
	var tags []entities.Tag
	var tag entities.Tag
	query := ("SELECT id, name FROM tb_tags")
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&tag.Id,
			&tag.Name,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	if len(tags) != 0 {
		return tags, nil
	}
	return tags, errors.New("tag is empty")

}
