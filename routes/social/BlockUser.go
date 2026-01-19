package social

import (
	"tbo_backend/objects"
	"tbo_backend/queries/social"
	"tbo_backend/utils"
	"tbo_backend/validations"

	"github.com/gofiber/fiber/v2"
)

func BlockUser(c *fiber.Ctx) error {

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
	valid := validations.ValidateBlockUser(&authObj)
	if !valid.Status {
		return c.Status(fiber.StatusOK).JSON(valid)
	}

	// get user id from context
	userId := c.Locals("userId").(string)

	// query on scylladb to unfollow user in following table and then put it in block list
	status := social.BlockUser(&authObj.UserId, &userId)
	if !status {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// return response
	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.UserBlockedSuccessfully,
	})

}
