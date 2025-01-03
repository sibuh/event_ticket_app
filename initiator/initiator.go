package initiator

import (
	"bus_ticket/internal/handler/ticket"
	"bus_ticket/internal/module/schedule"
	mtkt "bus_ticket/internal/module/ticket"
	"bus_ticket/internal/module/token"
	paymentintegration "bus_ticket/internal/platform/payment_integration"
	"context"

	huser "bus_ticket/internal/handler/user"
	"bus_ticket/internal/middleware"
	"bus_ticket/internal/utils/token/paseto"

	muser "bus_ticket/internal/module/user"
	"bus_ticket/internal/routing"

	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Initiate() {
	logger := InitLogger()
	InitConfig("config", logger)
	server := gin.Default()
	server.Use(middleware.Cors())
	v1 := server.Group("v1")
	fmt.Println("grouping:", v1)
	logger.Info("initiate database")
	queries := InitDB(
		// viper.GetString("dbconn"),
		fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")),
	)
	logger.Info("intiating storage layer")
	maker := paseto.NewPasetoMaker(viper.GetString("token.key"))
	mware := middleware.NewMiddleware(logger, maker, queries)
	token.Init(logger, queries, maker)
	sc := schedule.Init()
	module := NewModule(
		muser.Init(
			logger,
			queries,
			maker,
			viper.GetDuration("token.duration")),
		mtkt.Init(
			logger,
			paymentintegration.Init(logger, viper.GetString("payment.url")),
			queries, sc),
	)

	handler := InitHandler(
		huser.Init(logger, module.user),
		ticket.Init(logger, module.ticket),
	)
	routing.InitRouter(v1, handler.user, handler.ticket, mware)
	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		ReadTimeout: viper.GetDuration("server.read_time_out"),
		Handler:     server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	go func() {
		fmt.Println("server starting at ", viper.GetString("server.port"))
		srv.ListenAndServe()
	}()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	logger.Warn("sever is going to shut down %+V", srv.Shutdown(ctx))

}
