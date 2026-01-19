package profile

import (
	"tbo_backend/objects"
	"tbo_backend/queries/profile"
	"tbo_backend/utils"
	"tbo_backend/validations"

	"github.com/gofiber/fiber/v2"
)

func UpdateProfile(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// declare the auth object
	authObj := objects.ValidateProfileReq{}

	// Parse the JSON request body into the User struct
	if err := c.BodyParser(&authObj); err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// validate the payload
	valid := validations.ValidateProfile(&authObj)
	if !valid.Status {
		return c.Status(fiber.StatusOK).JSON(valid)
	}

	// get user id from locals
	userId := c.Locals("userId").(string)

	// update into database
	ok := profile.UpdateProfile(&authObj.Name, &authObj.Username, &userId)
	if !ok {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// return response
	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.ProfileUpdatedSuccess,
	})
}
