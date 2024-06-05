package routing

import (
	"event_ticket/internal/handler"
	"event_ticket/internal/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup, u handler.User, p handler.Payment, e handler.Event, t handler.Ticket, md middleware.Middleware) {
	routes := []Route{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: u.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: u.LoginUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/token",
			Handler: u.RefreshToken,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
		{
			Method:  http.MethodGet,
			Path:    "/pk",
			Handler: p.GetPublishableKey,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
		{
			Method:  http.MethodGet,
			Path:    "/cpi/:id",
			Handler: p.HandleCreatePaymentIntent,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
		{
			Method:  http.MethodPost,
			Path:    "/events",
			Handler: e.PostEvent,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
		{
			Method:  http.MethodGet,
			Path:    "/events",
			Handler: e.FetchEvents,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
		{
			Method:  http.MethodPost,
			Path:    "/webhook",
			Handler: p.PaymentWebhook,
		},
		{
			Method:  http.MethodGet,
			Path:    "/ticket/:intent_id",
			Handler: t.GetTicket,
			Mwares:  []gin.HandlerFunc{md.Authenticate()},
		},
	}
	RegisterRoutes(group, routes)

}
