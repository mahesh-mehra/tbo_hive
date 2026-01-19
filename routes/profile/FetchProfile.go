package profile

import (
	"tbo_backend/objects"
	"tbo_backend/queries/profile"
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func FetchProfile(c *fiber.Ctx) error {

	defer utils.HandleHttpPanic(c)

	// get user id from context
	userId := c.Locals("userId").(string)

	// fetch profile from database
	result, err := profile.FetchProfile(&userId)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// success response
	return c.Status(fiber.StatusOK).JSON(objects.ResponseWithData{
		Status: true,
		Data:   *result,
	})
}
