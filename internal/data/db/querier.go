// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetTicket(ctx context.Context, id uuid.UUID) (Ticket, error)
	GetTicketInfo(ctx context.Context, arg GetTicketInfoParams) (GetTicketInfoRow, error)
	GetTokenData(ctx context.Context) (GetTokenDataRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	StoreCheckoutSession(ctx context.Context, arg StoreCheckoutSessionParams) (Session, error)
	UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) (Ticket, error)
}

var _ Querier = (*Queries)(nil)
