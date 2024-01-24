package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
	"xyzApp/app/config"
	"xyzApp/app/helper"
	"xyzApp/app/model/auth"
	"xyzApp/app/model/dto"
)

func AuthMiddleware(cfg *config.AppConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// get token headeer
		header := ctx.Get("Authorization", "")
		tokenString := strings.Split(header, " ")
		token := tokenString[len(tokenString)-1]

		if header == "" || token == "" {
			statusCode := http.StatusUnauthorized
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    "token not set",
			})
		}

		// decode claims
		claims := auth.JwtClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Jwt.SecretKey), nil
		})

		// jika token tidak valid
		if err != nil {
			statusCode := http.StatusUnauthorized
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    "token isnt valid. you cant query this endpoint before login",
			})
		}

		// expired
		if time.Now().After(time.UnixMicro(claims.RegisteredClaims.ExpiresAt.Unix() * 1000000)) {
			statusCode := http.StatusUnauthorized
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    "token expired",
			})
		}

		ctx.Next()
		return nil
	}
}
