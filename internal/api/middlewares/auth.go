package middlewares

import (
	"log"
	"net/http"
	"tratodo/internal/libs/context"
	"tratodo/internal/libs/jwt"
	"tratodo/pkg/api"
)

// AuthMiddleware is a middleware that checks authentication before calling the http.Handler
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := inferClaimsJWT(r)
		if !ok {
			apiErr := api.NewApiError(http.StatusUnauthorized, "Unauthorized")
			api.WriteJSON(w, apiErr.StatusCode, apiErr)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.NewAuthContext(r.Context(), claims.ID)))
	})
}

func inferClaimsJWT(r *http.Request) (*jwt.AuthClaims, bool) {
	tokenStr := inferCookieJWT(r)
	if tokenStr == "" {
		return nil, false
	}

	claims, err := jwt.GetMapClaims(tokenStr)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	return claims, true
}

func inferCookieJWT(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}
