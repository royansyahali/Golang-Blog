package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	updaterequest "github.com/royansyahali/blog/payloads/request/updateRequest"
	"github.com/royansyahali/blog/repositories"
)

type UserRepositoryImp struct {
}

func NewUserRepository() repositories.UserRepository {
	return &UserRepositoryImp{}
}

func (u *UserRepositoryImp) InsertUser(ctx context.Context, tx *sql.Tx, reqUser *request.UserRequest) error {

	query := ("INSERT INTO tb_users (username, password, name, role, created_at, updated_at) VALUES (?,?,?,?,?,?)")
	_, err := tx.ExecContext(ctx, query,
		reqUser.Username,
		reqUser.Password,
		reqUser.Name,
		reqUser.Role,
		reqUser.CreateAt,
		reqUser.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImp) UpdateUser(ctx context.Context, tx *sql.Tx, reqUser *updaterequest.UserUpdate, id int) error {
	user, err := u.FindByIdUser(ctx, tx, id)
	if err != nil {
		return err
	}
	if reqUser.Name == "" {
		reqUser.Name = user.Name
	}
	if reqUser.Password == "" {
		reqUser.Password = user.Password
	}
	query := ("UPDATE tb_users SET password = ?, name = ?, updated_at = ? WHERE id = ?")
	_, err = tx.ExecContext(ctx, query,
		reqUser.Password,
		reqUser.Name,
		reqUser.UpdateAt,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryImp) DeleteUser(ctx context.Context, tx *sql.Tx, id int) error {
	_, err := u.FindByIdUser(ctx, tx, id)
	if err != nil {
		return err
	}
	query := ("DELETE FROM tb_users WHERE id = ?")
	_, err = tx.ExecContext(ctx, query,
		id,
	)
	if err != nil {
		return err
	}
	query = ("SELECT p.id FROM tb_users u JOIN tb_posts p ON u.id = p.user_id WHERE u.id = ?")
	_, err = tx.QueryContext(ctx, query,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImp) FindByIdUser(ctx context.Context, tx *sql.Tx, id int) (entities.User, error) {
	var user entities.User
	query := ("SELECT id, username, password, role, name, created_at, updated_at FROM tb_users WHERE id = ?")
	result, err := tx.QueryContext(ctx, query,
		id,
	)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
			&user.CreateAt,
			&user.UpdateAt,
		)
		if err != nil {
			return user, err
		}
		return user, err
	}

	return user, errors.New("user is not found in database")

}

func (u *UserRepositoryImp) GetAllUser(ctx context.Context, tx *sql.Tx) ([]entities.User, error) {
	var users []entities.User
	var user entities.User
	query := ("SELECT id, username, password,role ,name, created_at, updated_at FROM tb_users")
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
			&user.CreateAt,
			&user.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if len(users) != 0 {
		return users, nil
	}
	return users, errors.New("user is empty")

}
