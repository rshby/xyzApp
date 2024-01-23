package helper

import (
	"golang.org/x/crypto/bcrypt"
	"xyzApp/app/customError"
)

// function to hashed password
func HashPassword(password string) (string, error) {
	hashedPasswerd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", customError.NewBadRequestError(err.Error())
	}

	return string(hashedPasswerd), nil
}

// function to check password
func CheckPasword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, customError.NewBadRequestError("password not match")
	}

	// password match
	return true, nil
}
