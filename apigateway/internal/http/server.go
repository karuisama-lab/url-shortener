package http

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"url-shortener/apigateway/internal/http/handlers"
	"url-shortener/internal/config"
)

func RunServer(cfg *config.Config, lg *slog.Logger, deps *handlers.Deps) error {
	app := fiber.New()

	SetupRoutes(app, deps)

	lg.Info("starting HTTP server", "addr", cfg.HTTPServer.Address)

	if err := app.Listen(cfg.HTTPServer.Address); err != nil {
		return err
	}

	return nil
}
