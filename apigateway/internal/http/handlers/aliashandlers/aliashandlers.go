package aliashandlers

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"url-shortener/aliasservice/domain"
	"url-shortener/apigateway/internal/http/dto/aliasdto"
)

type AliasHandler struct {
	Service domain.AliasService
}

func NewAliasHandler(service domain.AliasService) *AliasHandler {
	return &AliasHandler{
		Service: service,
	}
}

func (handler *AliasHandler) SaveURL(ctx *fiber.Ctx) error {
	var req aliasdto.URLSaveRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	if req.URL == {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "url is required",
		})
		alias, err := handler.Service.SaveURL(ctx.Context(), req.URL)
		if err != nil {return err
			resp := aliasdto.URLSaveResponse{
				Message: "url saved",
				Alias: alias,
			}
			return ctx.Status(fiber.StatusCreated).JSON(resp)
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
