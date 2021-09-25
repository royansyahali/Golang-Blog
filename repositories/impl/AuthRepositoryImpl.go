package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/repositories"
)

type AuthRepositoryImp struct {
}

func NewAuthRepository() repositories.AuthRepository {
	return &AuthRepositoryImp{}
}

func (u *AuthRepositoryImp) Login(ctx context.Context, tx *sql.Tx, reqUser *entities.User) (entities.User, error) {
	var user entities.User
	query := ("SELECT username, password FROM tb_users WHERE username = ?")
	result, err := tx.QueryContext(ctx, query,
		reqUser.Username,
	)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(
			&user.Username,
			&user.Password,
		)
		if err != nil {
			return user, err
		}
		return user, err
	}

	return user, errors.New("user is not found in database")

}
