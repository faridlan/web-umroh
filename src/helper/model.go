package helper

import (
	"github.com/faridlan/web-umroh/src/model/domain"
	"github.com/faridlan/web-umroh/src/model/web"
)

func ToUserResponse(request domain.User) web.UserResponse {
	user := web.UserResponse{
		Id:       request.Id,
		Username: request.Username,
		Role:     request.Role,
	}

	return user
}

func ToUserReponses(request []domain.User) []web.UserResponse {
	userResponse := []web.UserResponse{}
	for _, user := range request {
		userResponse = append(userResponse, ToUserResponse(user))
	}

	return userResponse
}
