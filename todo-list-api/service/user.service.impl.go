package service

import (
	"context"
	"database/sql"
	"todo-list-api/helper"
	"todo-list-api/models/domain"
	"todo-list-api/models/web"
	"todo-list-api/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

// Create implements UserService.
func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

// Update implements UserService.
func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

// Delete implements UserService.
func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, user)
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

// FindByAll implements UserService.
func (service *UserServiceImpl) FindByAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindByAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, users := range users {
		userResponses = append(userResponses, helper.ToUserResponse(users))
	}
	return userResponses
}

// Login implements UserService.
func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	userFromDB := service.UserRepository.Login(ctx, tx, user)

	return helper.ToUserResponse(userFromDB)
}

// UpdateUsername implements UserService.
func (service *UserServiceImpl) UpdateUsername(ctx context.Context, request web.UserUpdateRequestUsername) (domain.User, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// validasi user
	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Username = request.Username

	_, err = service.UserRepository.UpdateUsername(ctx, tx, user)
	helper.PanicIfError(err)

	return user, nil
}
