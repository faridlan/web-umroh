package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-umroh/src/model/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}
