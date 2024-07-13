package ticket

import (
	"context"
	"database/sql"
	"errors"
	"event_ticket/internal/constant"
	"event_ticket/internal/data/db"
	"event_ticket/internal/model"
	"event_ticket/internal/storage"
	"net/http"

	"golang.org/x/exp/slog"
)

type ticket struct {
	logger *slog.Logger
	db     db.Querier
}

func Init(logger *slog.Logger, db db.Querier) storage.Ticket {
	return &ticket{
		logger: logger,
		db:     db,
	}
}
func (t *ticket) HoldTicket(ctx context.Context, req model.ReserveTicketRequest) (model.Ticket, error) {
	tkt, err := t.db.UpdateTicketStatus(ctx, db.UpdateTicketStatusParams{
		ID:     req.ID,
		Status: string(constant.Onhold),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := model.Error{
				ErrCode:   http.StatusNotFound,
				Message:   "the requested ticket is not found",
				RootError: err,
			}
			return model.Ticket{}, &newError
		}
		newError := model.Error{
			ErrCode:   http.StatusInternalServerError,
			Message:   "failed to update ticket status",
			RootError: err,
		}
		return model.Ticket{}, &newError
	}
	return model.Ticket{
		TripID:   tkt.TripID,
		BusNo:    tkt.BusNo,
		TicketNo: tkt.TicketNo,
		Status:   tkt.Status,
	}, nil
}
func (t *ticket) GetTicket(id string) (model.Ticket, error) {
	return model.Ticket{}, nil
}
func (t *ticket) UnholdTicket(tktNo, tripID int32) (model.Ticket, error) {
	return model.Ticket{}, nil
}
