// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Bus struct {
	ID uuid.UUID
}

type Session struct {
	ID            uuid.UUID
	TicketID      uuid.UUID
	PaymentStatus string
	PaymentUrl    string
	CancelUrl     sql.NullString
	Amount        float64
	CreatedAt     time.Time
}

type Ticket struct {
	ID     uuid.UUID
	TripID uuid.UUID
	BusID  uuid.UUID
	UserID uuid.UUID
	Status string
}

type Trip struct {
	ID uuid.UUID
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
