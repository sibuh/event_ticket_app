// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	AddEvent(ctx context.Context, arg AddEventParams) (Event, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	FetchEvent(ctx context.Context, id int32) (Event, error)
	FetchEvents(ctx context.Context) ([]Event, error)
	GetPayment(ctx context.Context, intentID string) (Payment, error)
	GetUser(ctx context.Context, username string) (User, error)
	RecordPayment(ctx context.Context, arg RecordPaymentParams) (Payment, error)
	UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) (Ticket, error)
	GetTicket(ctx context.Context,id string)(Ticket,error)
	StoreCheckoutSession(ctx context.Context, arg StoreCheckoutSessionParams) (Session, error)
	GetTicketStatus(ctx context.Context,sid string)(string, error)
}

var _ Querier = (*Queries)(nil)
