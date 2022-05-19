package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/sergejkoll/tg-example/pkg/adder_service"
	"github.com/sergejkoll/tg-example/pkg/config"
	"github.com/sergejkoll/tg-example/pkg/transport"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	log.Logger = config.Service().Logger()

	log.Log().Msg("service start")
	defer log.Log().Msg("service stop")

	svc := adder_service.NewService(log.Logger)

	services := []transport.Option{
		transport.Adder(transport.NewAdder(log.Logger, svc)),
	}

	srv := transport.New(log.Logger, services...).WithLog(log.Logger)

	srv.Fiber().Get("/api/healthcheck", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(map[string]string{"status": "Ok"})
	})

	go func() {
		log.Info().Str("bind", ":9000").Msg("listen on") // Move to config
		if err := srv.Fiber().Listen(":9000"); err != nil {
			log.Panic().Err(err).Stack().Msg("server error")
		}
	}()

	<-shutdown
}
