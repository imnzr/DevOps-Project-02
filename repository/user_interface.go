package repository

import (
	"context"
	"database/sql"

	"github.com/imnzr/DevOps-Project-02/models/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)

	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
