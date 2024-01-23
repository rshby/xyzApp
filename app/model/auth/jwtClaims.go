package auth

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	Email            string               `json:"email,omitempty"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims,omitempty"`
}

func (j *JwtClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return j.RegisteredClaims.GetExpirationTime()
}

func (j *JwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return j.RegisteredClaims.GetIssuedAt()
}

func (j *JwtClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return j.RegisteredClaims.GetNotBefore()
}

func (j *JwtClaims) GetIssuer() (string, error) {
	return j.RegisteredClaims.GetIssuer()
}

func (j *JwtClaims) GetSubject() (string, error) {
	return j.RegisteredClaims.GetSubject()
}

func (j *JwtClaims) GetAudience() (jwt.ClaimStrings, error) {
	return j.RegisteredClaims.GetAudience()
}
