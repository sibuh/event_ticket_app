package module

import (
	"context"
	"event_ticket/internal/model"
	"time"
)

type Ticket interface {
	ReserveTicket(ctx context.Context, req model.ReserveTicketRequest, scheduler func()) (model.Session, error)
	ScheduleOntimeoutProcess(ctx context.Context, delay time.Duration, url string)
}

type User interface {
	CreateUser(ctx context.Context, usr model.CreateUserRequest) (model.User, error)
	LoginUser(ctx context.Context, logReq model.LoginRequest) (string, error)
	RefreshToken(ctx context.Context, username string) (string, error)
}
type Event interface {
	PostEvent(ctx context.Context, postEvent model.Event) (model.Event, error)
	FetchEvents(ctx context.Context) ([]model.Event, error)
	FetchEvent(ctx context.Context, id int32) (model.Event, error)
}

type Payment interface {
	CreatePaymentIntent(ctx context.Context, userID, eventID int32) (string, error)
	GetPayment(ctx context.Context, intentID string) (model.Payment, error)
}
