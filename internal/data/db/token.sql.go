// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: token.sql

package db

import (
	"context"
)

const getTokenData = `-- name: GetTokenData :one
SELECT id, first_name, last_name, phone, email, username, password, created_at, updated_at, deleted_at
from users
`

func (q *Queries) GetTokenData(ctx context.Context) (User, error) {
	row := q.db.QueryRow(ctx, getTokenData)
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
