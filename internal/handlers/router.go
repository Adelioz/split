package handlers

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(l *zap.Logger) *mux.Router {
	router := mux.NewRouter()

	router.Use(
		RecoveryMiddleware(l),
		LoggingMiddleware(l),
	)

	router.Path("/hello").Handler(hello())
	return router
}
