package ticket

import (
	"context"
	"event_ticket/internal/model"
	"event_ticket/internal/module"
	"event_ticket/internal/platform"
	"event_ticket/internal/storage"
	"net/http"

	"golang.org/x/exp/slog"
)

type ticket struct {
	log           slog.Logger
	storageTicket storage.Ticket
	platform      platform.PaymentGatewayIntegrator
}
type TicketStatus string

const (
	Reserved TicketStatus = "Reserved"
	Free     TicketStatus = "Free"
	Onhold   TicketStatus = "Onhold"
)

func Init(log slog.Logger, tkt storage.Ticket, platform platform.PaymentGatewayIntegrator) module.Ticket {
	return &ticket{
		log:           log,
		storageTicket: tkt,
		platform:      platform,
	}
}
func (t *ticket) ReserveTicket(ctx context.Context, tktNo, tripId int32) (string, error) {
	tkt, err := t.storageTicket.GetTicket(tktNo, tripId)
	if err != nil {
		return "", err
	}
	if tkt.Status == string(Reserved) {
		newError := model.Error{
			ErrCode:   http.StatusBadRequest,
			Message:   "ticket is already reserved please try to reserve free ticket",
			RootError: nil,
		}
		return "", &newError
	}

	if tkt.Status == string(Free) {
		newError := model.Error{
			ErrCode:   http.StatusBadRequest,
			Message:   "ticket is onhold please try later",
			RootError: nil,
		}
		return "", &newError
	}

	tkt, err = t.storageTicket.ReserveTicket(tktNo, tripId)

	if err != nil {
		return "", err
	}
	url, err := t.platform.CreateCheckoutSession(ctx, tkt)
	if err != nil {
		return "", err
	}
	return url, nil
}
