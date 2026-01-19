package social

import (
	"tbo_backend/objects"
	"tbo_backend/queries/social"
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func FetchBlockedUserList(c *fiber.Ctx) error {

	defer utils.HandleHttpPanic(c)

	// get user id from context
	userId := c.Locals("userId").(string)

	// query on scylladb to get blocked user list
	blockedUserList := social.FetchBlockedUserList(&userId)

	return c.Status(fiber.StatusOK).JSON(objects.ResponseWithBlockedUserList{
		Status: true,
		Data:   blockedUserList,
	})
}
