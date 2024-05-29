// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: events.sql

package db

import (
	"context"
	"time"
)

const addEvent = `-- name: AddEvent :one
INSERT INTO events (title,description,user_id,start_date,end_date) VALUES ($1,$2,$3,$4,$5)
RETURNING id, title, description, user_id, start_date, end_date, created_at, updated_at, deleted_at
`

type AddEventParams struct {
	Title       string
	Description string
	UserID      int32
	StartDate   time.Time
	EndDate     time.Time
}

func (q *Queries) AddEvent(ctx context.Context, arg AddEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, addEvent,
		arg.Title,
		arg.Description,
		arg.UserID,
		arg.StartDate,
		arg.EndDate,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getEvent = `-- name: GetEvent :one
SELECT id, title, description, user_id, start_date, end_date, created_at, updated_at, deleted_at FROM events
`

func (q *Queries) GetEvent(ctx context.Context) (Event, error) {
	row := q.db.QueryRow(ctx, getEvent)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}