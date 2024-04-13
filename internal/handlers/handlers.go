package handlers

import (
	"net/http"

	"github.com/Adelioz/split/internal/models"
	"github.com/Adelioz/split/pkg/logging"
)

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := models.Model{
			Name: "Adel",
		}
		logging.S(r.Context()).Debug("hell yeah")
		WriteJSONResponse(r.Context(), w, http.StatusOK, response)
	})
}
