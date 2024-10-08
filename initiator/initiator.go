package initiator

import (
	"context"
	"event_ticket/internal/handler/ticket"
	"event_ticket/internal/module/schedule"
	mtkt "event_ticket/internal/module/ticket"
	paymentintegration "event_ticket/internal/platform/payment_integration"

	"log"

	huser "event_ticket/internal/handler/user"
	"event_ticket/internal/middleware"
	"event_ticket/internal/utils/token/paseto"

	muser "event_ticket/internal/module/user"
	"event_ticket/internal/routing"

	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"github.com/spf13/viper"
)

func Initiate() {
	logger := InitLogger()
	InitConfig("config", logger)
	server := gin.Default()
	server.Use(middleware.Cors())
	v1 := server.Group("v1")
	logger.Info("initiate database")
	fmt.Println("connection string", fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME")))
	queries := InitDB(fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	logger.Info("intiating storage layer")
	maker := paseto.NewPasetoMaker(viper.GetString("token.key"), viper.GetDuration("token.duration"))
	mware := middleware.NewMiddleware(logger, maker, queries)
	sc := schedule.Init()
	module := NewModule(
		muser.Init(logger, queries, maker),
		mtkt.Init(logger,
			paymentintegration.Init(logger, viper.GetString("payment.url")),
			queries, sc),
	)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	handler := InitHandler(
		huser.Init(logger, module.user),
		ticket.Init(logger, module.ticket),
	)
	routing.InitRouter(v1, handler.user, handler.ticket, mware)
	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		ReadTimeout: viper.GetDuration("server.read_time_out") * time.Second,
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
