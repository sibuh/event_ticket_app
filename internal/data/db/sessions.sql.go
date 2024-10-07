// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: sessions.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const storeCheckoutSession = `-- name: StoreCheckoutSession :one
INSERT INTO sessions (id,ticket_id,payment_status,payment_url,cancel_url,amount,created_at)VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, ticket_id, payment_status, payment_url, cancel_url, amount, created_at
`

type StoreCheckoutSessionParams struct {
	ID            uuid.UUID
	TicketID      uuid.UUID
	PaymentStatus string
	PaymentUrl    string
	CancelUrl     sql.NullString
	Amount        float64
	CreatedAt     time.Time
}

func (q *Queries) StoreCheckoutSession(ctx context.Context, arg StoreCheckoutSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, storeCheckoutSession,
		arg.ID,
		arg.TicketID,
		arg.PaymentStatus,
		arg.PaymentUrl,
		arg.CancelUrl,
		arg.Amount,
		arg.CreatedAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.TicketID,
		&i.PaymentStatus,
		&i.PaymentUrl,
		&i.CancelUrl,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
