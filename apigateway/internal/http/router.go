package http

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/apigateway/internal/http/handlers"
)

func SetupRoutes(app *fiber.App, h *handlers.Deps) *fiber.App {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	alias := v1.Group("/alias")
	alias.Post("/", h.Alias.SaveURL)
	alias.Get("/:alias", h.Alias.GetURL)

	return app
}
