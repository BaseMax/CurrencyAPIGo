package middlewares

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

// Render put render.Render into context
func RenderMiddleware(r *render.Render) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), 3, r)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
