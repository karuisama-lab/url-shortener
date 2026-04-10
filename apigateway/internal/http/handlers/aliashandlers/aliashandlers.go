package aliashandlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"url-shortener/aliasservice/domain"
	"url-shortener/apigateway/internal/http/dto/aliasdto"
	"url-shortener/apigateway/internal/transport/clients/aliasclient"
)

type AliasHandler struct {
	Logger  *slog.Logger
	Service domain.AliasInterface
}

func NewAliasHandler(logger *slog.Logger, client *aliasclient.Client) *AliasHandler {
	return &AliasHandler{
		Logger: logger,
		Client: client,
	}
}

func (handler *AliasHandler) SaveURL(ctx *fiber.Ctx) error {
	var req aliasdto.URLSaveRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	alias, err := handler.Service.SaveURL(ctx.Context(), req.URL)
	if err != nil {
		return err
	}
	resp := aliasdto.URLSaveResponse{
		"url saved",
		alias,
	}
	return ctx.Status(201).JSON(resp)

}

func (handler *AliasHandler) GetURL(ctx *fiber.Ctx) error {
	alias := ctx.Params("alias")
	if alias == "" {
		return ctx.Status(400).JSON(map[string]string{
			"message": "alias is required",
		})
	}
	return ctx.Status(200).JSON(aliasdto.URLGetResponse{
		URL:     "https.example.com",
		Message: "url is found",
		Alias:   alias,
	})
}
