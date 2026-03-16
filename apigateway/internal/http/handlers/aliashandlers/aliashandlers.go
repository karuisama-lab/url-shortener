package aliashandlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"url-shortener/apigateway/internal/http/dto/aliasdto"
	"url-shortener/apigateway/internal/transport/clients/aliasclient"
)

type AliasHandler struct {
	Logger *slog.Logger
	Client *aliasclient.Client
}

func NewAliasHandler(logger *slog.Logger, client *aliasclient.Client) *AliasHandler {
	return &AliasHandler{
		Logger: logger,
		Client: client,
	}
}

func (h *AliasHandler) SaveURL(c *fiber.Ctx) error {
	var req aliasdto.URLSaveRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.UserContext()

	resp, err := h.Client.SaveURL(req, ctx)
	if err != nil {
		return err
	}

	c.JSON(aliasdto.URLSaveResponse{
		Code:    fiber.StatusOK,
		Message: resp.Message,
	})

	return nil
}
