package web

import "github.com/faridlan/web-umroh/src/model/domain"

type UserResponse struct {
	Id       int
	Username string
	Role     domain.Role
}
