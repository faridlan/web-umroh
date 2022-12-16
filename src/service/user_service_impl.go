package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-umroh/src/helper"
	"github.com/faridlan/web-umroh/src/model/domain"
	"github.com/faridlan/web-umroh/src/model/web"
	"github.com/faridlan/web-umroh/src/repository"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateReq) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}

	user = service.UserRepo.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateReq) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Username = request.Username
	user.Password = request.Password
	user.Role = request.Role

	user = service.UserRepo.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	service.UserRepo.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepo.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToUserReponses(users)
}
