package web

import "github.com/faridlan/web-umroh/src/model/domain"

type UserUpdateReq struct {
	Id       int
	Username string
	Password string
	Role     domain.Role
}
