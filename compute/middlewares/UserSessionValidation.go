package middlewares

import (
	"fmt"
	"strings"
	"tbo_backend/objects"
	"tbo_backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserSessionValidation(c *fiber.Ctx) error {

	// handling exception
	defer utils.HandleHttpPanic(c)

	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(objects.Response{
			Status: false,
			Msg:    objects.InvalidAuthorization,
		})
	}

	tokenString := strings.Split(authorization[0], " ")
	if len(tokenString) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(objects.Response{
			Status: false,
			Msg:    objects.InvalidAuthorization,
		})
	}

	token := tokenString[1]

	userClaims := &objects.UserClaims{}
	_, err := jwt.ParseWithClaims(token, userClaims, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is HMAC and not something else
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(objects.ConfigObj.SecretKey), nil
	})

	if err != nil {
		println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(objects.Response{
			Status: false,
			Msg:    objects.InvalidAuthorization,
		})
	}

	c.Locals("userId", userClaims.Contact)
	c.Locals("name", userClaims.Name)

	return c.Next()
}
