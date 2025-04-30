package helper

import (
	"todo-list-api/models/domain"
	"todo-list-api/models/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
