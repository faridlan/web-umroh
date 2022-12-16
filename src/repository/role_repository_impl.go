package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/web-umroh/src/helper"
	"github.com/faridlan/web-umroh/src/model/domain"
)

type RoleRepositoryImpl struct {
}

func (repository RoleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {

	SQL := "insert into role(name) values (?)"
	res, err := tx.ExecContext(ctx, SQL, role.Name)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)

	return role
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update role SET name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Name, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error) {
	SQL := "select id, name from role where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}

	if rows.Next() {
		err := rows.Scan(&role.Id, role.Name)
		helper.PanicIfError(err)

		return role, nil
	} else {
		return role, errors.New("role not found")
	}
}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, name from role"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	roles := []domain.Role{}

	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, role.Name)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}

	return roles
}
