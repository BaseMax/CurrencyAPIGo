package middlewares

import (
	"context"
	"net/http"

	"github.com/itsjoniur/currency/internal/utils"
)

// CurrencyLogger put utils.Logger into context
func CurrencyLogger(logger *utils.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), 2, logger)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
