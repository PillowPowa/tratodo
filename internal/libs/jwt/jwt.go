package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	ID  int64 `json:"id"`
	Exp int64 `json:"exp"`
	jwt.MapClaims
}

func NewToken(payload int64, ttl time.Duration) (string, error) {
	claims := AuthClaims{
		ID:  payload,
		Exp: time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("jwt.NewToken: %w", err)
	}

	return t, nil
}

func GetMapClaims(tokenStr string) (*AuthClaims, error) {
	const op = "jwt.GetMapClaims"

	token, err := ValidateJWT(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("%s: Invalid token", op)
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		return nil, fmt.Errorf("%s: Unexpected claims type", op)
	}

	return claims, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	const op = "jwt.ValidateJWT"
	secret := os.Getenv("JWT_SECRET")

	return jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: Unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
