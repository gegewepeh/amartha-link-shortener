package v1

import (
	"context"
	"errors"
	"link-shortener/internal/app/link_shortener/models"
	"link-shortener/internal/pkg/utils"
)

func (pool *Pool) FindUser(ctx context.Context, username string, password string) (*models.User, *utils.AppError) {
	var userId string
	var passwordData string

	query := `SELECT users.id, users.username, users.password FROM users
	WHERE username = $1`
	err := pool.db.QueryRow(ctx, query, username).Scan(&userId, &username, &passwordData)
	if err != nil {
		return nil, models.WrapError("User", "InvalidParameters", errors.New("Invalid Username or Password"), nil)
	}
	if password != passwordData {
		return nil, models.WrapError("User", "InvalidParameters", errors.New("Invalid Username or Password"), nil)
	}
	user := models.User{
		ID: userId,
	}

	return &user, nil
}

func (pool *Pool) CreateUser(ctx context.Context, username string, password string) (*models.User, *utils.AppError) {
	query := `INSERT INTO users(username, password, "createdAt", "updatedAt")
	VALUES ($1, $2, NOW(), NOW())
	RETURNING users.username`

	err := pool.db.QueryRow(ctx, query, username, password).Scan(&username)
	if err != nil {
		return nil, models.WrapError("Link", "InternalServerError", err, nil)
	}
	user := models.User{
		Username: username,
	}

	return &user, nil
}
