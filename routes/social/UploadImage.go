package social

import (
	"tbo_backend/objects"
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {

	defer utils.HandleHttpPanic(c)

	// // read multipart form data
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
	// 		Status: false,
	// 		Msg:    "Invalid request",
	// 	})
	// }

	// // get user id from context
	// userId := c.Locals("userId").(string)

	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.AccountDeleteSuccessfully,
	})
}
