package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/web-umroh/src/helper"
	"github.com/faridlan/web-umroh/src/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user (username, password, role_ind) VALUE (?,?,?,?)"
	res, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)

	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "Update user Set username = ?, role_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "Delete from user where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)

}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, username, role_id where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Role.Id)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, username, role_id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Role.Id)
		helper.PanicIfError(err)

		users = append(users, user)
	}
	return users
}
