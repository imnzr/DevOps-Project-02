package repository

import (
	"context"
	"database/sql"
	"todo-list-api/models/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.User

	// Login user by email
	Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
