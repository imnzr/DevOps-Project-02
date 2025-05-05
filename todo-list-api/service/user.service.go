package service

import (
	"context"
	"todo-list-api/models/domain"
	"todo-list-api/models/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindByAll(ctx context.Context) []web.UserResponse

	// Login find by email
	Login(ctx context.Context, email web.LoginRequest) web.UserResponse
	UpdateUsername(ctx context.Context, request web.UserUpdateRequestUsername) (domain.User, error)
}
