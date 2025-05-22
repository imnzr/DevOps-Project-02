package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/imnzr/DevOps-Project-02/helper"
	"github.com/imnzr/DevOps-Project-02/models/domain"
	"github.com/imnzr/DevOps-Project-02/models/web"
	"github.com/imnzr/DevOps-Project-02/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpplementation struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpplementation{
		UserRepository: userRepository,
		DB:             db,
	}
}

// Create implements UserService.
func (service *UserServiceImpplementation) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing password: %v", err)
	}
	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	savedUser, err := service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		log.Printf("error saving user: %v", err)
	}

	return web.UserResponse{
		Id:       savedUser.Id,
		Username: savedUser.Username,
		Email:    savedUser.Email,
	}
}

// Delete implements UserService.
func (service *UserServiceImpplementation) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		log.Printf("error deleting user: %v", err)
	}

	service.UserRepository.Delete(ctx, tx, user)
}

// FindAll implements UserService.
func (service *UserServiceImpplementation) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	users := service.UserRepository.FindByAll(ctx, tx)
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		})
	}
	return userResponses
}

// FindById implements UserService.
func (service *UserServiceImpplementation) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		log.Printf("error finding user: %v", err)
	}

	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

// Login implements UserService.
func (service *UserServiceImpplementation) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(errors.New("user not found"))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Printf("error comparing password: %v", err)
	}
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// Update implements UserService.
func (service *UserServiceImpplementation) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("error begin transaction: %v", err)
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		log.Printf("user not found: %v", err)
	}
	user.Username = request.Username
	user.Email = request.Email

	// Hash password baru
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("error hashing password: %v", err)
		}
		user.Password = string(hashedPassword)
	}

	userUpdated, err := service.UserRepository.Update(ctx, tx, user)
	if err != nil {
		log.Printf("error updating user: %v", err)
	}

	return web.UserResponse{
		Id:       userUpdated.Id,
		Username: userUpdated.Username,
		Email:    userUpdated.Email,
	}
}
