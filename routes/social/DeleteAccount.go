package social

import (
	"tbo_backend/objects"
	"tbo_backend/queries/social"

	"github.com/gofiber/fiber/v2"
)

func DeleteAccount(c *fiber.Ctx) error {

	// update into user table to delete account
	// get user id from context
	userId := c.Locals("userId").(string)

	status := social.DeleteAccount(&userId)
	if !status {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.AccountDeleteSuccessfully,
	})
}
