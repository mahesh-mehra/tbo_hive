package profile

import (
	"tbo_backend/objects"
	"tbo_backend/queries/profile"
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadProfilePhoto(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// 1. Get the file from the form input (key: "photo")
	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.PhotoMandatory,
		})
	}

	// 2. Validate File Size (Example: Limit to 5MB)
	if file.Size > 5*1024*1024 {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.ImageTooLarge,
		})
	}

	userId := c.Locals("userId").(string)
	imageName := userId + "_profile.png"

	// 3. Save the file to a specific directory
	dst := objects.ConfigObj.LocalPath.ProfilePhotos + "/" + imageName
	if err := c.SaveFile(file, dst); err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// 4. update into database
	ok := profile.UpdateProfilePhoto(&userId, &imageName)
	if !ok {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// success response
	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.ProfilePhotoUpdatedSuccessfully,
	})
}
