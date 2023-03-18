package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(ttl time.Duration, payload interface{}, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token_string, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return token_string, nil
}

func VerifyJwt(token string, signedJwtKey string) (jwt.MapClaims, error) {
	tkn, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJwtKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)

	if !ok || !tkn.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims,  nil
}
