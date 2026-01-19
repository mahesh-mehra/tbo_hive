package utils

import (
	"fmt"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ztrue/tracerr"
	"tbo_backend/objects"
)

// HandlePanic handle panic function, it recovers the application from exception
func HandlePanic() {

	//if any error is thrown, then it pushes the exception into kafka
	if err := recover(); err != nil {
		text := tracerr.Sprint(err.(error))
		fmt.Println(text)
	}
}

func HandleHttpPanic(c *fiber.Ctx) {
	if err := recover(); err != nil {
		_ = c.JSON(&objects.Response{
			Status: false,
			Msg:    fmt.Sprint(err),
		})
	}
}
