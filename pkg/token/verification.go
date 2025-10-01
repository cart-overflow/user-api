package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyAccessToken(
	token string,
	secret string,
	timeFunc func() time.Time,
) (*TokenClaims, error) {
	parserOpts := []jwt.ParserOption{
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	}
	if timeFunc != nil {
		parserOpts = append(parserOpts, jwt.WithTimeFunc(timeFunc))
	}

	jwt, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, parserOpts...)
	if err != nil {
		return nil, err
	}

	return jwt.Claims.(*TokenClaims), nil
}
