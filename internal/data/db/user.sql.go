// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name,last_name,phone,email,username,password) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, first_name, last_name, phone, email, username, password, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Username  string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.Email,
		arg.Username,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, phone, email, username, password, created_at, updated_at, deleted_at FROM users WHERE username=$1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
