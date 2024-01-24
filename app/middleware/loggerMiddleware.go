package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/config"
)

func LoggerMiddleware(cfg config.IConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Capture request body
		var bodyMap map[string]any
		json.Unmarshal(ctx.Body(), &bodyMap)
		fmt.Println(bodyMap)
		ctx.Next()

		var responseMap map[string]any
		json.Unmarshal(ctx.Response().Body(), &responseMap)
		fmt.Println(responseMap)

		fmt.Println(string(ctx.Request().URI().Path()))    // endpoint
		fmt.Println(ctx.Response().StatusCode())           // statuscode
		fmt.Println(string(ctx.Request().Header.Method())) // method
		return nil
	}
}
