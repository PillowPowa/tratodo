package middlewares

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CORS(allowOrigin string) mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{allowOrigin}),
		handlers.AllowedHeaders([]string{"Content-Type", "Cookie"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PATCH", "PUT", "OPTIONS"}),
	)
}
