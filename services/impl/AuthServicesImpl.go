package impl

import (
	"context"
	"database/sql"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/repositories"
	"github.com/royansyahali/blog/securities"
	"github.com/royansyahali/blog/services"
)

type AuthServiceImpl struct {
	AuthRepository repositories.AuthRepository
	DB             *sql.DB
}

func NewAuthService(u repositories.AuthRepository, db *sql.DB) services.AuthService {
	return &AuthServiceImpl{AuthRepository: u, DB: db}
}

func (u *AuthServiceImpl) Login(ctx context.Context, reqUser *entities.User) error {
	var user entities.User
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	user, err = u.AuthRepository.Login(ctx, tx, reqUser)
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
	err = securities.VerifyPassword(user.Password, reqUser.Password)
	if err != nil {
		return err
	}
	return nil
}
