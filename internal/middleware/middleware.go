package middleware

import (
	"event_ticket/internal/data/db"
	"event_ticket/internal/utils/token"
	"event_ticket/internal/utils/token/paseto"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

const authType string = "Bearer"

type Middleware struct {
	logger *slog.Logger
	maker  token.TokenMaker
	db.Querier
}

func NewMiddleware(logger *slog.Logger, maker token.TokenMaker, q db.Querier) Middleware {
	return Middleware{
		logger:  logger,
		maker:   maker,
		Querier: q,
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			m.logger.Info("authorization header is empty")
			ctx.AbortWithStatus(401)
			return
		}
		authSlice := strings.Split(auth, " ")
		if authSlice[0] != authType {
			m.logger.Info(fmt.Sprintf("invalide authorization type want:%s got:%s", authType, authSlice[0]))
			ctx.AbortWithStatus(401)
			return
		}
		tokenMaker := paseto.NewPasetoMaker(viper.GetString("token.key"), viper.GetDuration("token.duration"))
		payload, err := tokenMaker.VerifyToken(authSlice[1])
		if err != nil {
			m.logger.Info("failed to verify token", err)
			ctx.AbortWithStatus(401)
			return
		}
		if !payload.Valid() {
			m.logger.Info("token is not valid please refresh token")
			ctx.AbortWithStatus(401)
			return
		}
		usr, err := m.Querier.GetUser(ctx, payload.Username)
		if err != nil {
			m.logger.Info("user does not exist", err)
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Set("user", usr)
	}
}
