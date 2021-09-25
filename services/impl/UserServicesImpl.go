package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	updaterequest "github.com/royansyahali/blog/payloads/request/updateRequest"
	"github.com/royansyahali/blog/repositories"
	"github.com/royansyahali/blog/securities"
	"github.com/royansyahali/blog/services"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validasi       *validator.Validate
}

func NewUserService(u repositories.UserRepository, db *sql.DB, v *validator.Validate) services.UserService {
	return &UserServiceImpl{UserRepository: u, DB: db, Validasi: v}
}

func (u *UserServiceImpl) InsertUser(ctx context.Context, reqUser *request.UserRequest) error {
	err := u.Validasi.Struct(reqUser)
	if err != nil {
		return err
	}
	if reqUser.Role != "admin" && reqUser.Role != "guest" {
		return errors.New("role user is invalid")
	}
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	reqUser.Password, err = securities.HashPassword(reqUser.Password)
	if err != nil {
		return err
	}
	err = u.UserRepository.InsertUser(ctx, tx, reqUser)
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

func (u *UserServiceImpl) UpdateUser(ctx context.Context, reqUser *updaterequest.UserUpdate, id int) error {
	err := u.Validasi.Struct(reqUser)
	if err != nil {
		return err
	}
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	reqUser.Password, err = securities.HashPassword(reqUser.Password)
	if err != nil {
		return err
	}
	err = u.UserRepository.UpdateUser(ctx, tx, reqUser, id)
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

func (u *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	err = u.UserRepository.DeleteUser(ctx, tx, id)
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

func (u *UserServiceImpl) FindByIdUser(ctx context.Context, id int) (entities.User, error) {
	var user entities.User
	tx, err := u.DB.Begin()
	if err != nil {
		return user, err
	}
	user, err = u.UserRepository.FindByIdUser(ctx, tx, id)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return user, errRollback
		}
		return user, err
	}
	err = tx.Commit()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserServiceImpl) GetAllUser(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	users, err = u.UserRepository.GetAllUser(ctx, tx)
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
	return users, nil
}
