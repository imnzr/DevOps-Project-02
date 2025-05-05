package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-list-api/helper"
	"todo-list-api/models/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "UPDATE `user` SET username = ?, email = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user

}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	query := "DELETE FROM user WHERE id = ?"
	_, err := tx.QueryContext(ctx, query, user.Id)
	helper.PanicIfError(err)
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT id, username FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// FindByAll implements UserRepository.
func (u *UserRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, username, email FROM user"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

// Login implements UserRepository.
func (u *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "SELECT id, email, username, password FROM user WHERE email = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, query, user.Email)

	var dbUser domain.User

	err := row.Scan(&dbUser.Id, &dbUser.Email, &dbUser.Username, &dbUser.Password)
	helper.PanicIfError(err)

	if dbUser.Password != user.Password {
		panic("password salah")
	}
	return dbUser
}

// UpdateUsername implements UserRepository.
func (u *UserRepositoryImpl) UpdateUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "UPDATE user SET username = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Username, user.Id)

	return user, err
}
