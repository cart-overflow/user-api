package token

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	UserId            string `json:"user_id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	RegisterTimestamp int64  `json:"register_timestamp"`
	Type              string `json:"type"`
	jwt.RegisteredClaims
}
