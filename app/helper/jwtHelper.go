package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"xyzApp/app/config"
	"xyzApp/app/model/auth"
)

func GenerateToken(cfg *config.AppConfig, email string) (string, error) {
	claims := &auth.JwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "xyzApp",
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(1 * time.Hour)),
		}}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString([]byte(cfg.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	// sucess generate token
	return tokenString, nil
}
