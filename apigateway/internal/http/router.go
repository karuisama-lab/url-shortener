package http

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/internal/transport/http/handlers"
)

package http

import (
"github.com/gofiber/fiber/v2"
"urlShortener/internal/transport/http/handlers"
)

func SetupRoutes(app *fiber.App, h *handlers.Handlers) *fiber.App{
	api := app.Group("/api")
	v1 := api.Group("/v1")

	alias := v1.Group("/alias")
	alias.Post("/", h.URL.SaveURL)
	alias.Get("/:alias", h.URL.GetURL)

	return app
}
