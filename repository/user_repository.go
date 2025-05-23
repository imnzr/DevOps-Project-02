package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/DevOps-Project-02/helper"
	"github.com/imnzr/DevOps-Project-02/models/domain"
)

type UserRepositoryImplementation struct{}

// FindByEmail implements UserRepository.

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (u *UserRepositoryImplementation) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT id, username, email, password FROM `user` WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.HandleQueryError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		helper.HandleErrorRows(err)

		return user, nil
	} else {
		return domain.User{}, fmt.Errorf("user with email %s not found", email)
	}
}

// Delete implements UserRepository.
func (u *UserRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := "DELETE FROM `user` WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Id)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %d: %w", user.Id, err)
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected after deleting user with ID %d: %w", user.Id, err)
	}
	if RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", user.Id)
	}

	return nil
}

// FindByAll implements UserRepository.
func (u *UserRepositoryImplementation) FindByAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, username, email FROM `user`"
	rows, err := tx.QueryContext(ctx, query)
	helper.HandleQueryError(err)

	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		users = append(users, user)
	}
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT id, username, email FROM `user` WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.HandleQueryError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.HandleErrorRows(err)
		return user, nil
	} else {
		return domain.User{}, fmt.Errorf("user with ID %d not found", userId)
	}
}

// Login implements UserRepository.
func (u *UserRepositoryImplementation) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "SELECT id, username, email FROM `user` WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, query, user.Email, user.Password)
	helper.HandleQueryError(err)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.HandleErrorRows(err)
		return user, nil
	} else {
		return domain.User{}, fmt.Errorf("invalid username or password")
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helper.HandleQueryError(err)

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	user.Id = int(lastInsertId)
	return user, nil
}

// Update implements UserRepository.
func (u *UserRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "UPDATE `user` SET username = ?, email = ?, password = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id)
	helper.HandleQueryError(err)

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return domain.User{}, fmt.Errorf("no user found with ID %d", user.Id)
	}
	return user, nil
}
