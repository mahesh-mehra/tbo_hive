package social

import (
	"tbo_backend/objects"
	"tbo_backend/queries/social"
	"tbo_backend/utils"
	"tbo_backend/validations"

	"github.com/gofiber/fiber/v2"
)

func Follow(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// declare the auth object
	authObj := objects.FollowReq{}

	// Parse the JSON request body into the User struct
	if err := c.BodyParser(&authObj); err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// valdiate the requests
	valid := validations.ValidateFollow(&authObj)
	if !valid.Status {
		return c.Status(fiber.StatusOK).JSON(valid)
	}

	// get user id from context
	userId := c.Locals("userId").(string)

	// check if the user is already following the user if not then follow
	status := social.FollowRequest(&authObj.UserId, &userId)
	if !status {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.FollowingSuccessfully,
	})
}
