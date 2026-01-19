package auth

import (
	"tbo_backend/objects"
	"tbo_backend/services"
	"tbo_backend/utils"
	"tbo_backend/validations"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// declare the auth object
	authObj := objects.AuthReq{}

	// Parse the JSON request body into the User struct
	if err := c.BodyParser(&authObj); err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// validate the request
	validateUser := validations.ValidateUser(&authObj)
	if !validateUser.Status {
		return c.Status(fiber.StatusOK).JSON(validateUser)
	}

	// Use Service
	success, err := h.authService.Login(authObj.Mobile)
	if err != nil || !success {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// return response
	return c.Status(fiber.StatusOK).JSON(objects.Response{
		Status: true,
		Msg:    objects.OtpSentSuccessfully,
	})
}

func (h *AuthHandler) ValidateOtp(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	// declare the auth object
	authObj := objects.ValidateOtpReq{}

	// Parse the JSON request body into the User struct
	if err := c.BodyParser(&authObj); err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// validate the request
	valid := validations.ValidateMobileVerify(&authObj)
	if !valid.Status {
		return c.Status(fiber.StatusOK).JSON(valid)
	}

	// Use Service
	name, ok, err := h.authService.ValidateOtp(authObj.Mobile, authObj.Otp)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}
	if !ok {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.InvalidOtp,
		})
	}

	// generate token
	token := utils.GenerateJWTToken(&authObj.Mobile, &name)
	if token == "" {
		return c.Status(fiber.StatusOK).JSON(objects.Response{
			Status: false,
			Msg:    objects.DefaultResp,
		})
	}

	// return response
	return c.Status(fiber.StatusOK).JSON(objects.AuthResponse{
		Status: true,
		Token:  token,
	})
}
