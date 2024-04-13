package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(ctx context.Context, w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Error: %s", err)
		// logging.S(ctx).Errorf("Failed to write response: %w", err)
	}
}
