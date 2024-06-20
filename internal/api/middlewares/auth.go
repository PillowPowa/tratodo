package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"tratodo/internal/libs/jwt"
	"tratodo/pkg/api"
)

func WithAuth(fn api.HandlerFunc) api.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		tokenStr := inferCookieJWT(r)
		if tokenStr == "" {
			return api.NewApiError(http.StatusUnauthorized, "Unauthorized")
		}

		claims, err := jwt.GetMapClaims(tokenStr)
		if err != nil {
			log.Println(err)
			return api.NewApiError(http.StatusUnauthorized, "Unauthorized")
		}

		fmt.Println("[Auth Middleware] Found user ID:", claims.ID)
		return fn(w, r)
	}
}

func inferCookieJWT(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}
