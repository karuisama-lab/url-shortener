package httperr

import "github.com/gofiber/fiber/v2"

func WriteHTTPError(c fiber.Ctx, err error) error {
	return c.JSON(err)
}
