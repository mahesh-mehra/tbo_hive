package profile

import (
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateEmail(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// return response
	return c.SendString("Login")
}
