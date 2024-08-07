package middlewares

import (
	"net/http"
	"tratodo/internal/libs/context"
	"tratodo/internal/libs/jwt"
	"tratodo/pkg/api"
)

// AuthMiddleware is a middleware that checks authentication before calling the http.Handler
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		api.MakeHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			claims, ok := inferClaimsJWT(r)
			if !ok {
				return api.NewApiError(http.StatusUnauthorized, "Unauthorized")
			}
			next.ServeHTTP(w, r.WithContext(context.NewAuthContext(r.Context(), claims.ID)))
			return nil
		}),
	)
}

func inferClaimsJWT(r *http.Request) (*jwt.AuthClaims, bool) {
	tokenStr := inferCookieJWT(r)
	if tokenStr == "" {
		return nil, false
	}

	claims, err := jwt.GetMapClaims(tokenStr)
	if err != nil {
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
