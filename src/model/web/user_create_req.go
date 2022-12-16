package web

import "github.com/faridlan/web-umroh/src/model/domain"

type UserCreateReq struct {
	Username string
	Password string
	Role     domain.Role
}
