package handlers

import (
	"net/http"

	"github.com/Adelioz/split/pkg/logging"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// LoggingMiddleware populates request context with logger. Could also be a good place to log
// request data if needed.
func LoggingMiddleware(l *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := logging.WithLogger(r.Context(), l)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// RecoveryMiddleware catches panics and responds with 500.
func RecoveryMiddleware(l *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			func() {
				defer func() {
					if i := recover(); i != nil {
						l.Sugar().Errorf("Recovered panic: %v.", i)
						w.WriteHeader(http.StatusInternalServerError)
					}
				}()
				next.ServeHTTP(w, r)
			}()
		})
	}
}
