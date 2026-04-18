package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"url-shortener/aliasservice/handlers"
	"url-shortener/aliasservice/storage/postgres"
	"url-shortener/aliasservice/usecase"
)

func main() {
	app := fiber.New()

	repo := postgres.NewPostgres()
	service := usecase.NewService(repo)
	aliasHandler := handler.NewAliasHandler(service)

	app.Post("/url", aliasHandler.SaveURL)

	log.Fatal(app.Listen(":8080"))
}
