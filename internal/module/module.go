package module

import (
	"context"
	"event_ticket/internal/model"

	"github.com/signintech/gopdf"
)

type Ticket interface {
	// CreateCheckoutSession(c *gin.Context, user model.User) error
	// UpdatePaymentStatus(status, sid string) (db.User, error)
	GeneratePDFTicket(intentID string) (*gopdf.GoPdf, error)
}

// type Sms interface {
// 	SendSms(user db.User, wg *sync.WaitGroup) error
// }

// type Email interface {
// 	SendEmail(user db.User, attachmentPath string, wg *sync.WaitGroup) error
// }

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
